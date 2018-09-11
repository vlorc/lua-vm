package tcp

import (
	"context"
	"github.com/vlorc/lua-vm/base"
	vmnet "github.com/vlorc/lua-vm/net"
)

func NewTCPFactory(driver vmnet.NetDriver, opt ...func(*TCPFactory)) *TCPFactory {
	o := &TCPFactory{
		driver:         driver,
		context:        context.Background(),
		connectTimeout: base.Duration(5000),
		listenTimeout:  base.Duration(5000),
	}
	for _, v := range opt {
		v(o)
	}
	return o
}
