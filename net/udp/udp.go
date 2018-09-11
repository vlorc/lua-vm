package udp

import (
	"context"
	"net"
	"strconv"
	"time"
	vmnet "github.com/vlorc/lua-vm/net"
)

type UDPConn struct {
	conn net.PacketConn
	remote net.Addr
	readTimeout time.Duration
	writeTimeout time.Duration
}

type UDPFactory struct {
	driver vmnet.NetDriver
}

func(f *UDPFactory)Listen(addr string,args ...int) (*UDPConn,error) {
	return f.__listen("udp",addr,args...)
}
func(f *UDPFactory)Listen4(addr string,args ...int) (*UDPConn,error) {
	return f.__listen("udp4",addr,args...)
}
func(f *UDPFactory)Listen6(addr string,args ...int) (*UDPConn,error) {
	return f.__listen("udp6",addr,args...)
}
func(f *UDPFactory)Connect(addr string,args ...int) (*UDPConn,error) {
	return f.__listen("udp",addr,args...)
}
func(f *UDPFactory)Connect4(addr string,args ...int) (*UDPConn,error) {
	return f.__listen("udp4",addr,args...)
}
func(f *UDPFactory)Connect6(addr string,args ...int) (*UDPConn,error) {
	return f.__listen("udp6",addr,args...)
}
func(f *UDPFactory)__listen(network,addr string,args ...int) (*UDPConn,error) {
	if len(args) > 0 && args[0] >= 0{
		addr = net.JoinHostPort(addr,strconv.Itoa(args[0]))
	}
	conn,err := f.driver.Packet(context.Background(),network,addr)
	if nil != err {
		return nil,err
	}
	return &UDPConn{conn: conn},nil
}
func(f *UDPFactory)__connect(network,addr string,args ...int) (net.PacketConn,error) {
	return nil,nil
}
