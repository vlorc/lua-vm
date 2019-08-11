package driver

import (
	"context"
	vmnet "github.com/vlorc/lua-vm/net"
	"net"
)

type DirectDriver struct {
	dialer net.Dialer
	listen net.ListenConfig
}

func NewDirectDriver(dialer net.Dialer, listen net.ListenConfig) vmnet.NetDriver {
	return &DirectDriver{
		dialer: dialer,
		listen: listen,
	}
}

func (d *DirectDriver) Dial(ctx context.Context, network, addr string) (net.Conn, error) {
	return d.dialer.DialContext(ctx, network, addr)
}

func (d *DirectDriver) Listen(ctx context.Context, network, addr string) (net.Listener, error) {
	return d.listen.Listen(context.Background(), network, addr)
}

func (d *DirectDriver) Packet(ctx context.Context, network, addr string) (net.PacketConn, error) {
	return d.listen.ListenPacket(ctx, network, addr)
}
