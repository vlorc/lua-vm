package net

import (
	"context"
	"net"
)

type NetDriver interface {
	Dial(ctx context.Context, network, addr string) (net.Conn, error)
	Packet(ctx context.Context, network, address string) (net.PacketConn, error)
	Listen(ctx context.Context, network, addr string) (net.Listener, error)
}

type DialFunc func(context.Context, string, string) (net.Conn, error)

func (f DialFunc) Dial(network, addr string) (net.Conn, error) {
	return f(context.Background(), network, addr)
}
