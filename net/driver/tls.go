package driver

import (
	"context"
	"crypto/tls"
	vmnet "github.com/vlorc/lua-vm/net"
	"net"
	"strings"
)

type tlsDriver struct {
	vmnet.NetDriver
	config *tls.Config
}

func NewTLSDriver(config *tls.Config, parent vmnet.NetDriver) (vmnet.NetDriver, error) {
	if nil == config {
		config = &tls.Config{InsecureSkipVerify: true}
	} else {
		config = config.Clone()
	}

	return &tlsDriver{NetDriver: parent, config: config}, nil
}

func (h *tlsDriver) Dial(ctx context.Context, network, addr string) (conn net.Conn, err error) {
	if conn, err = h.NetDriver.Dial(ctx, network, addr); nil != err {
		return nil, err
	}
	if strings.HasSuffix(network, "tcp") {
		conn, err = h.__upgrade(ctx, conn)
	}
	return conn, err
}

func (tp *tlsDriver) __upgrade(ctx context.Context, raw net.Conn) (net.Conn, error) {
	conn := tls.Client(raw, tp.config)
	done := make(chan error, 1)
	go func() {
		done <- conn.Handshake()
	}()
	var err error
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-done:
	}
	if close(done); nil != err {
		conn.Close()
		conn = nil
	}
	return conn, err
}
