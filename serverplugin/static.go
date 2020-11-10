package serverplugin

import (
	"net"
)

// StaticRegisterPlugin implements static registry.
type StaticRegisterPlugin struct {
}

// Start starts to connect etcd cluster
func (p *StaticRegisterPlugin) Start() error {
	return nil
}

// HandleConnAccept handles connections from clients
func (p *StaticRegisterPlugin) HandleConnAccept(conn net.Conn) (net.Conn, bool) {
	return conn, true
}

// Register handles registering event.
// this service is registered at BASE/serviceName/thisIpAddress node
func (p *StaticRegisterPlugin) Register(name string, rcvr interface{}, metadata string) (err error) {
	return
}

// Deregister 注销服务
func (p *StaticRegisterPlugin) Deregister(name string) (err error) {
	return
}
