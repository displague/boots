package job

import (
	"net"
	"time"

	"github.com/packethost/pkg/log"
	"github.com/pkg/errors"
	"github.com/tinkerbell/boots/conf"
	"github.com/tinkerbell/boots/dhcp"
	"github.com/tinkerbell/boots/packet"
	tw "github.com/tinkerbell/tink/protos/workflow"
)

var client *packet.Client

// SetClient sets the client used to interact with the api.
func SetClient(c *packet.Client) {
	client = c
}

// Job this comment is useless
type Job struct {
	log.Logger

	mac net.HardwareAddr
	ip  net.IP

	start time.Time

	mode Mode
	dhcp dhcp.Config

	hardware packet.Hardware
	instance *packet.Instance
}

// hasActiveWorkflow fetches workflows for a given hwID and returns
// the status true if there is a pending (active) workflow for hwID
// hwID is the hardware/worker ID corresponding to the MAC
func hasActiveWorkflow(wcl *tw.WorkflowContextList) (bool, error) {
	for _, wf := range (*wcl).WorkflowContexts {
		if wf.CurrentActionState == tw.ActionState_ACTION_PENDING {
			joblog.With("workflowID", wf.WorkflowId).Info("found active workflow for hardware")
			return true, nil
		}
	}
	return false, nil
}

// CreateFromDHCP looks up hardware using the MAC from cacher to create a job
func CreateFromDHCP(mac net.HardwareAddr, giaddr net.IP, circuitID string) (Job, error) {
	j := Job{
		mac:   mac,
		start: time.Now(),
	}

	d, err := discoverHardwareFromDHCP(mac, giaddr, circuitID)
	if err != nil {
		return Job{}, errors.WithMessage(err, "discover from dhcp message")
	}

	err = j.setup(d)
	if err != nil {
		return Job{}, err
	}
	return j, nil
}

// CreateFromRemoteAddr looks up hardware using the IP from cacher to create a job
func CreateFromRemoteAddr(ip string) (Job, error) {
	host, _, err := net.SplitHostPort(ip)
	if err != nil {
		return Job{}, errors.Wrap(err, "splitting host:ip")
	}
	return CreateFromIP(net.ParseIP(host))
}

// CreateFromIP looksup hardware using the IP from cacher to create a job
func CreateFromIP(ip net.IP) (Job, error) {
	j := Job{
		ip:    ip,
		start: time.Now(),
	}

	joblog.With("ip", ip).Info("discovering from ip")
	d, err := discoverHardwareFromIP(ip)
	if err != nil {
		return Job{}, errors.WithMessage(err, "discovering from ip address")
	}
	mac := d.GetMAC(ip)
	if mac.String() == packet.ZeroMAC.String() {
		joblog.With("ip", ip).Fatal(errors.New("somehow got a zero mac"))
	}
	j.mac = mac

	err = j.setup(d)
	if err != nil {
		return Job{}, err
	}

	hd := d.Hardware()
	hwID := hd.HardwareID()

	joblog.With("hardwareID", hwID).Info("fetching workflows for hardware")
	wcl, err := client.GetWorkflowsFromTink(hwID)
	if err != nil {
		return Job{}, err
	}

	activeWorkflows, err := hasActiveWorkflow(wcl)
	if err != nil {
		return Job{}, err
	}

	if !activeWorkflows {
		return Job{}, errors.Errorf("no active workflow found for hardware %s", hwID)
	}

	return j, nil
}

// MarkDeviceActive marks the device active
func (j Job) MarkDeviceActive() {
	if id := j.InstanceID(); id != "" {
		if err := client.PostInstancePhoneHome(id); err != nil {
			j.Error(err)
		}
	}
}

func (j *Job) setup(d packet.Discovery) error {
	dh := d.Hardware()

	j.Logger = joblog.With("mac", j.mac, "hardware.id", dh.HardwareID())

	// mac is needed to find the hostname for DiscoveryCacher
	d.SetMAC(j.mac)

	// (kdeng3849) is this necessary?
	j.hardware = d.Hardware()

	// (kdeng3849) how can we remove this?
	j.instance = d.Instance()
	if j.instance == nil {
		j.instance = &packet.Instance{}
	} else {
		j.Logger = j.Logger.With("instance.id", j.InstanceID())
	}

	ip := d.GetIP(j.mac)
	if ip.Address == nil {
		return errors.New("could not find IP address")
	}
	j.dhcp.Setup(ip.Address, ip.Netmask, ip.Gateway)
	j.dhcp.SetLeaseTime(d.LeaseTime(j.mac))
	j.dhcp.SetDHCPServer(conf.PublicIPv4) // used for the unicast DHCPREQUEST
	j.dhcp.SetDNSServers(d.DnsServers(j.mac))

	hostname, err := d.Hostname()
	if err != nil {
		return err
	}
	if hostname != "" {
		j.dhcp.SetHostname(hostname)
	}

	return nil
}
