package tcp

import (
	"context"
	"net"
	"time"
	vmnet"github.com/vlorc/lua-vm/net"
)

type TCPConn struct {
	conn net.Conn
	readTimeout time.Duration
	writeTimeout time.Duration
}

type TCPListener struct {
	listen net.Listener
	readTimeout time.Duration
	writeTimeout time.Duration
}

type TCPFactory struct {
	driver vmnet.NetDriver
	context context.Context
	resolveTimeout time.Duration
	listenTimeout time.Duration
	connectTimeout time.Duration
	readTimeout time.Duration
	writeTimeout time.Duration
}
