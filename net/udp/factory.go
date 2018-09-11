package udp

import "github.com/vlorc/lua-vm/net"

func NewUDPFactory(driver net.NetDriver) *UDPFactory {
	return &UDPFactory{
		driver: driver,
	}
}
