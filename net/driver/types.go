package driver

import (
	vmnet "github.com/vlorc/lua-vm/net"
	"net/url"
)

type Factory func(uri *url.URL, parent, factory vmnet.NetDriver) (vmnet.NetDriver, error)
