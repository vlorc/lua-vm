package udp

import (
	"context"
	vmnet "github.com/vlorc/lua-vm/net"
	"net"
	"strconv"
	"time"
)

type UDPConn struct {
	conn         net.PacketConn
	remote       *net.UDPAddr
	readTimeout  time.Duration
	writeTimeout time.Duration
}

type UDPFactory struct {
	driver vmnet.NetDriver
}

func (f *UDPFactory) Listen(addr string, args ...int) (*UDPConn, error) {
	return f.__listen("udp", addr, args...)
}
func (f *UDPFactory) Listen4(addr string, args ...int) (*UDPConn, error) {
	return f.__listen("udp4", addr, args...)
}
func (f *UDPFactory) Listen6(addr string, args ...int) (*UDPConn, error) {
	return f.__listen("udp6", addr, args...)
}
func (f *UDPFactory) Connect(addr string, args ...int) (*UDPConn, error) {
	return f.__connect("udp", addr, args...)
}
func (f *UDPFactory) Connect4(addr string, args ...int) (*UDPConn, error) {
	return f.__connect("udp4", addr, args...)
}
func (f *UDPFactory) Connect6(addr string, args ...int) (*UDPConn, error) {
	return f.__connect("udp6", addr, args...)
}
func (f *UDPFactory) Resolve(addr string, args ...int) (net.Addr, error) {
	return __resolve("udp", addr, args...)
}
func (f *UDPFactory) Resolve4(addr string, args ...int) (net.Addr, error) {
	return __resolve("udp4", addr, args...)
}
func (f *UDPFactory) Resolve6(addr string, args ...int) (net.Addr, error) {
	return __resolve("udp6", addr, args...)
}
func (f *UDPFactory) __listen(network, addr string, args ...int) (*UDPConn, error) {
	if len(args) > 0 && args[0] >= 0 {
		addr = net.JoinHostPort(addr, strconv.Itoa(args[0]))
	}
	conn, err := f.driver.Packet(context.Background(), network, addr)
	if nil != err {
		return nil, err
	}
	return &UDPConn{conn: conn}, nil
}
func (f *UDPFactory) __connect(network, addr string, args ...int) (*UDPConn, error) {
	remote, err := __resolve(network, addr, args...)
	if nil != err {
		return nil, err
	}
	conn, err := f.driver.Packet(context.Background(), network, ":0")
	if nil != err {
		return nil, err
	}

	return &UDPConn{conn: conn, remote: remote.(*net.UDPAddr)}, nil
}
func __resolve(network, addr string, args ...int) (net.Addr, error) {
	if len(args) > 0 && args[0] >= 0 {
		addr = net.JoinHostPort(addr, strconv.Itoa(args[0]))
	}
	return net.ResolveUDPAddr(network, addr)
}
