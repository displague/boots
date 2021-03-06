// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tinkerbell/tink/protos/workflow (interfaces: WorkflowSvcClient)

// Package mocks is a generated GoMock package.
package mock_workflow

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	workflow "github.com/tinkerbell/tink/protos/workflow"
	grpc "google.golang.org/grpc"
)

// MockWorkflowSvcClient is a mock of WorkflowSvcClient interface
type MockWorkflowSvcClient struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowSvcClientMockRecorder
}

// MockWorkflowSvcClientMockRecorder is the mock recorder for MockWorkflowSvcClient
type MockWorkflowSvcClientMockRecorder struct {
	mock *MockWorkflowSvcClient
}

// NewMockWorkflowSvcClient creates a new mock instance
func NewMockWorkflowSvcClient(ctrl *gomock.Controller) *MockWorkflowSvcClient {
	mock := &MockWorkflowSvcClient{ctrl: ctrl}
	mock.recorder = &MockWorkflowSvcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkflowSvcClient) EXPECT() *MockWorkflowSvcClientMockRecorder {
	return m.recorder
}

// CreateWorkflow mocks base method
func (m *MockWorkflowSvcClient) CreateWorkflow(arg0 context.Context, arg1 *workflow.CreateRequest, arg2 ...grpc.CallOption) (*workflow.CreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWorkflow", varargs...)
	ret0, _ := ret[0].(*workflow.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkflow indicates an expected call of CreateWorkflow
func (mr *MockWorkflowSvcClientMockRecorder) CreateWorkflow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkflow", reflect.TypeOf((*MockWorkflowSvcClient)(nil).CreateWorkflow), varargs...)
}

// DeleteWorkflow mocks base method
func (m *MockWorkflowSvcClient) DeleteWorkflow(arg0 context.Context, arg1 *workflow.GetRequest, arg2 ...grpc.CallOption) (*workflow.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWorkflow", varargs...)
	ret0, _ := ret[0].(*workflow.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWorkflow indicates an expected call of DeleteWorkflow
func (mr *MockWorkflowSvcClientMockRecorder) DeleteWorkflow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflow", reflect.TypeOf((*MockWorkflowSvcClient)(nil).DeleteWorkflow), varargs...)
}

// GetWorkflow mocks base method
func (m *MockWorkflowSvcClient) GetWorkflow(arg0 context.Context, arg1 *workflow.GetRequest, arg2 ...grpc.CallOption) (*workflow.Workflow, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkflow", varargs...)
	ret0, _ := ret[0].(*workflow.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflow indicates an expected call of GetWorkflow
func (mr *MockWorkflowSvcClientMockRecorder) GetWorkflow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflow", reflect.TypeOf((*MockWorkflowSvcClient)(nil).GetWorkflow), varargs...)
}

// GetWorkflowActions mocks base method
func (m *MockWorkflowSvcClient) GetWorkflowActions(arg0 context.Context, arg1 *workflow.WorkflowActionsRequest, arg2 ...grpc.CallOption) (*workflow.WorkflowActionList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkflowActions", varargs...)
	ret0, _ := ret[0].(*workflow.WorkflowActionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowActions indicates an expected call of GetWorkflowActions
func (mr *MockWorkflowSvcClientMockRecorder) GetWorkflowActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowActions", reflect.TypeOf((*MockWorkflowSvcClient)(nil).GetWorkflowActions), varargs...)
}

// GetWorkflowContext mocks base method
func (m *MockWorkflowSvcClient) GetWorkflowContext(arg0 context.Context, arg1 *workflow.GetRequest, arg2 ...grpc.CallOption) (*workflow.WorkflowContext, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkflowContext", varargs...)
	ret0, _ := ret[0].(*workflow.WorkflowContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowContext indicates an expected call of GetWorkflowContext
func (mr *MockWorkflowSvcClientMockRecorder) GetWorkflowContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowContext", reflect.TypeOf((*MockWorkflowSvcClient)(nil).GetWorkflowContext), varargs...)
}

// GetWorkflowContexts mocks base method
func (m *MockWorkflowSvcClient) GetWorkflowContexts(arg0 context.Context, arg1 *workflow.WorkflowContextRequest, arg2 ...grpc.CallOption) (*workflow.WorkflowContextList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkflowContexts", varargs...)
	ret0, _ := ret[0].(*workflow.WorkflowContextList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowContexts indicates an expected call of GetWorkflowContexts
func (mr *MockWorkflowSvcClientMockRecorder) GetWorkflowContexts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowContexts", reflect.TypeOf((*MockWorkflowSvcClient)(nil).GetWorkflowContexts), varargs...)
}

// GetWorkflowData mocks base method
func (m *MockWorkflowSvcClient) GetWorkflowData(arg0 context.Context, arg1 *workflow.GetWorkflowDataRequest, arg2 ...grpc.CallOption) (*workflow.GetWorkflowDataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkflowData", varargs...)
	ret0, _ := ret[0].(*workflow.GetWorkflowDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowData indicates an expected call of GetWorkflowData
func (mr *MockWorkflowSvcClientMockRecorder) GetWorkflowData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowData", reflect.TypeOf((*MockWorkflowSvcClient)(nil).GetWorkflowData), varargs...)
}

// GetWorkflowDataVersion mocks base method
func (m *MockWorkflowSvcClient) GetWorkflowDataVersion(arg0 context.Context, arg1 *workflow.GetWorkflowDataRequest, arg2 ...grpc.CallOption) (*workflow.GetWorkflowDataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkflowDataVersion", varargs...)
	ret0, _ := ret[0].(*workflow.GetWorkflowDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowDataVersion indicates an expected call of GetWorkflowDataVersion
func (mr *MockWorkflowSvcClientMockRecorder) GetWorkflowDataVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowDataVersion", reflect.TypeOf((*MockWorkflowSvcClient)(nil).GetWorkflowDataVersion), varargs...)
}

// GetWorkflowMetadata mocks base method
func (m *MockWorkflowSvcClient) GetWorkflowMetadata(arg0 context.Context, arg1 *workflow.GetWorkflowDataRequest, arg2 ...grpc.CallOption) (*workflow.GetWorkflowDataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkflowMetadata", varargs...)
	ret0, _ := ret[0].(*workflow.GetWorkflowDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowMetadata indicates an expected call of GetWorkflowMetadata
func (mr *MockWorkflowSvcClientMockRecorder) GetWorkflowMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowMetadata", reflect.TypeOf((*MockWorkflowSvcClient)(nil).GetWorkflowMetadata), varargs...)
}

// ListWorkflows mocks base method
func (m *MockWorkflowSvcClient) ListWorkflows(arg0 context.Context, arg1 *workflow.Empty, arg2 ...grpc.CallOption) (workflow.WorkflowSvc_ListWorkflowsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWorkflows", varargs...)
	ret0, _ := ret[0].(workflow.WorkflowSvc_ListWorkflowsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkflows indicates an expected call of ListWorkflows
func (mr *MockWorkflowSvcClientMockRecorder) ListWorkflows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkflows", reflect.TypeOf((*MockWorkflowSvcClient)(nil).ListWorkflows), varargs...)
}

// ReportActionStatus mocks base method
func (m *MockWorkflowSvcClient) ReportActionStatus(arg0 context.Context, arg1 *workflow.WorkflowActionStatus, arg2 ...grpc.CallOption) (*workflow.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReportActionStatus", varargs...)
	ret0, _ := ret[0].(*workflow.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportActionStatus indicates an expected call of ReportActionStatus
func (mr *MockWorkflowSvcClientMockRecorder) ReportActionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportActionStatus", reflect.TypeOf((*MockWorkflowSvcClient)(nil).ReportActionStatus), varargs...)
}

// ShowWorkflowEvents mocks base method
func (m *MockWorkflowSvcClient) ShowWorkflowEvents(arg0 context.Context, arg1 *workflow.GetRequest, arg2 ...grpc.CallOption) (workflow.WorkflowSvc_ShowWorkflowEventsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ShowWorkflowEvents", varargs...)
	ret0, _ := ret[0].(workflow.WorkflowSvc_ShowWorkflowEventsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowWorkflowEvents indicates an expected call of ShowWorkflowEvents
func (mr *MockWorkflowSvcClientMockRecorder) ShowWorkflowEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowWorkflowEvents", reflect.TypeOf((*MockWorkflowSvcClient)(nil).ShowWorkflowEvents), varargs...)
}

// UpdateWorkflowData mocks base method
func (m *MockWorkflowSvcClient) UpdateWorkflowData(arg0 context.Context, arg1 *workflow.UpdateWorkflowDataRequest, arg2 ...grpc.CallOption) (*workflow.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateWorkflowData", varargs...)
	ret0, _ := ret[0].(*workflow.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWorkflowData indicates an expected call of UpdateWorkflowData
func (mr *MockWorkflowSvcClientMockRecorder) UpdateWorkflowData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowData", reflect.TypeOf((*MockWorkflowSvcClient)(nil).UpdateWorkflowData), varargs...)
}
