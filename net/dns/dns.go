package dns

import (
	"context"
	"github.com/vlorc/lua-vm/base"
	vmnet "github.com/vlorc/lua-vm/net"
	"net"
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

func (f *DNSFactory) Lookup(host string, args ...int) ([]string, error) {
	ctx := context.Background()
	if len(args) > 0 && args[0] > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, base.Duration(args[0]))
		defer cancel()
	}
	addr, err := f.resolver.LookupIPAddr(ctx, host)
	if err != nil {
		return nil, err
	}
	ips := make([]string, len(addr))
	for i, a := range addr {
		ips[i] = a.IP.String()
	}
	return ips, nil
}

func (f *DNSFactory) LookupAddr(host string, args ...int) ([]string, error) {
	ctx := context.Background()
	if len(args) > 0 && args[0] > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, base.Duration(args[0]))
		defer cancel()
	}
	return f.resolver.LookupAddr(ctx, host)
}

func (f *DNSFactory) LookupHost(host string, args ...int) ([]string, error) {
	ctx := context.Background()
	if len(args) > 0 && args[0] > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, base.Duration(args[0]))
		defer cancel()
	}
	return f.resolver.LookupHost(ctx, host)
}

func (f *DNSFactory) LookupCNAME(host string, args ...int) (string, error) {
	ctx := context.Background()
	if len(args) > 0 && args[0] > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, base.Duration(args[0]))
		defer cancel()
	}
	return f.resolver.LookupCNAME(ctx, host)
}

func (f *DNSFactory) LookupMX(host string, args ...int) ([]*net.MX, error) {
	ctx := context.Background()
	if len(args) > 0 && args[0] > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, base.Duration(args[0]))
		defer cancel()
	}
	return f.resolver.LookupMX(ctx, host)
}

func (f *DNSFactory) LookupNS(host string, args ...int) ([]*net.NS, error) {
	ctx := context.Background()
	if len(args) > 0 && args[0] > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, base.Duration(args[0]))
		defer cancel()
	}
	return f.resolver.LookupNS(ctx, host)
}

func (f *DNSFactory) LookupTXT(host string, args ...int) ([]string, error) {
	ctx := context.Background()
	if len(args) > 0 && args[0] > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, base.Duration(args[0]))
		defer cancel()
	}
	return f.resolver.LookupTXT(ctx, host)
}
