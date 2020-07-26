FROM --platform=$BUILDPLATFORM golang:1.14-alpine as builder
ARG TARGETARCH
WORKDIR /src
COPY . /src/
RUN GOARCH=$TARGETARCH CGO_ENABLED=0 go build -v -ldflags="-X main.GitRev=$(shell git rev-parse --short HEAD)"

FROM --platform=$BUILDPLATFORM alpine:3.10 as boots
EXPOSE 67 69 80
RUN apk add --update --upgrade --no-cache ca-certificates socat
COPY --from=builder /src/boots /
ENTRYPOINT ["/boots"]
