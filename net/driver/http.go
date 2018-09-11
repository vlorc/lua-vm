package driver

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	vmnet "github.com/vlorc/lua-vm/net"
	"net"
	"net/http"
	"net/url"
	"strings"
	"unsafe"
)

type HttpDriver struct {
	vmnet.NetDriver
	host   string
	format string
	https  bool
	config *tls.Config
}

func NewHttpDriver(rawurl string, parent vmnet.NetDriver) (vmnet.NetDriver, error) {
	uri, err := url.Parse(rawurl)
	if nil != err {
		return nil, err
	}
	return __newHttpDriver(uri, parent)
}

func __newHttpDriver(uri *url.URL, parent vmnet.NetDriver) (vmnet.NetDriver, error) {
	driver := &HttpDriver{
		NetDriver: parent,
		host:      uri.Host,
		https:     "https" == uri.Scheme,
		config:    &tls.Config{},
		format:    "CONNECT %s HTTP/1.1\nHost: %s\n\n",
	}
	if nil != uri.User {
		password, _ := uri.User.Password()
		driver.format = fmt.Sprintf(
			"CONNECT %%s HTTP/1.1\nHost: %%s\nAuthorization: Basic %s\n\n",
			base64.StdEncoding.EncodeToString([]byte(uri.User.Username()+":"+password)))

	}
	return driver, nil
}

func (h *HttpDriver) Dial(ctx context.Context, network, addr string) (net.Conn, error) {
	if strings.HasSuffix(network, "tcp") {
		return h.__dial(ctx, network, addr)
	}
	return h.NetDriver.Dial(ctx, network, addr)
}

func (h *HttpDriver) __dialProxy(ctx context.Context, network, addr string) (conn net.Conn, err error) {
	rawConn, err := h.NetDriver.Dial(ctx, network, addr)
	if !h.https {
		conn = rawConn
		return
	}
	tlsConn := tls.Client(conn, h.config)
	errChannel := make(chan error, 1)
	go func() {
		errChannel <- tlsConn.Handshake()
	}()
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-errChannel:
	}
	if nil != err {
		rawConn.Close()
		return nil, err
	}
	return tlsConn, nil
}

func (h *HttpDriver) __dial(ctx context.Context, network, addr string) (net.Conn, error) {
	conn, err := h.__dialProxy(ctx, network, addr)
	var body = fmt.Sprintf(h.format, addr, addr)
	if _, err = conn.Write(*(*[]byte)(unsafe.Pointer(&body))); nil != err {
		conn.Close()
		return nil, err
	}
	resp, err := http.ReadResponse(bufio.NewReader(conn), nil)
	resp.Body.Close()
	if err != nil {
		conn.Close()
		return nil, err
	}
	if resp.StatusCode != 200 {
		conn.Close()
		err = fmt.Errorf("Connect server using proxy error, StatusCode [%d]", resp.StatusCode)
		return nil, err
	}
	return conn, nil
}
