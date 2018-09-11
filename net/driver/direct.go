package driver

import (
	"context"
	"net"
)

type DirectDriver struct {}

func(DirectDriver)Dial(ctx context.Context,network, addr string) (net.Conn, error){
	d := net.Dialer{}
	return d.DialContext(ctx,network,addr)
}

func(DirectDriver)Listen(ctx context.Context,network, addr string) (net.Listener, error){
	return net.Listen(network,addr)
}

func(DirectDriver)Packet(ctx context.Context,network, address string) (net.PacketConn, error) {
	return net.ListenPacket(network,address)
}
