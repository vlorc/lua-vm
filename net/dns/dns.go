package dns

import (
	"context"
	"net"
	"github.com/vlorc/lua-vm/base"
	vmnet "github.com/vlorc/lua-vm/net"
)

type DNSFactory struct {
	resolver net.Resolver
}

func NewDNSFactory(driver vmnet.NetDriver) *DNSFactory {
	return &DNSFactory{
		resolver: net.Resolver{
			Dial: driver.Dial,
		},
	}
}

func (f *DNSFactory) Lookup(host string,args ...int) ([]net.IPAddr,error){
	ctx := context.Background()
	if len(args) > 0 && args[0] > 0{
		var cancel context.CancelFunc
		ctx,cancel = context.WithTimeout(ctx,base.Duration(args[0]))
		defer cancel()
	}
	return f.resolver.LookupIPAddr(ctx,host)
}
