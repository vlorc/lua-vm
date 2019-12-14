package driver

import (
	"fmt"
	vmnet "github.com/vlorc/lua-vm/net"
	"net/url"
)

var mapping = map[string]Factory{
	"http":   __newHttpDriver,
	"https":  __newHttpsDriver,
	"sock5":  __newSock5Driver,
	"sock5s": __newSock5SslDriver,
}

func Register(driver string, factory Factory) {
	mapping[driver] = factory
}

func NewProxyDriver(rawurl string, parent vmnet.NetDriver) (vmnet.NetDriver, error) {
	uri, err := url.Parse(rawurl)
	if nil != err {
		return nil, err
	}
	factory, ok := mapping[uri.Scheme]
	if !ok {
		return nil, fmt.Errorf("scheme '%s' not support", uri.Scheme)
	}
	return factory(uri, parent, parent)
}
