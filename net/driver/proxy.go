package driver

import (
	"fmt"
	"net/url"
	vmnet "github.com/vlorc/lua-vm/net"
)

func NewProxyDriver(rawurl string, parent vmnet.NetDriver) (vmnet.NetDriver, error) {
	uri, err := url.Parse(rawurl)
	if nil != err {
		return nil, err
	}
	switch uri.Scheme {
	case "http","https":
		return __newHttpDriver(uri,parent)
	case "sock5":
		return __newSock5Driver(uri,parent)
	}
	return nil,fmt.Errorf("scheme '%s' not support",uri.Scheme)
}
