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

type httpDriver struct {
	vmnet.NetDriver
	factory vmnet.NetDriver
	host    string
	format  string
}

func NewHttpDriver(rawurl string, parent vmnet.NetDriver) (vmnet.NetDriver, error) {
	uri, err := url.Parse(rawurl)
	if nil != err {
		return nil, err
	}
	return __newHttpDriver(uri, parent, parent)
}

func NewHttpsDriverWithConfig(rawurl string, parent vmnet.NetDriver, config *tls.Config) (vmnet.NetDriver, error) {
	uri, err := url.Parse(rawurl)
	if nil != err {
		return nil, err
	}
	factory, err := NewTLSDriver(config, parent)
	if nil != err {
		return nil, err
	}
	return __newHttpDriver(uri, parent, factory)
}

func NewHttpsDriver(rawurl string, parent vmnet.NetDriver) (vmnet.NetDriver, error) {
	return NewHttpsDriverWithConfig(rawurl, parent, nil)
}

func __newHttpsDriver(uri *url.URL, parent, factory vmnet.NetDriver) (vmnet.NetDriver, error) {
	temp, err := NewTLSDriver(nil, factory)
	if nil != err {
		return nil, err
	}
	return __newHttpDriver(uri, parent, temp)
}

func __newHttpDriver(uri *url.URL, parent, factory vmnet.NetDriver) (vmnet.NetDriver, error) {
	driver := &httpDriver{
		NetDriver: parent,
		factory:   factory,
		host:      uri.Host,
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

func (h *httpDriver) Dial(ctx context.Context, network, addr string) (net.Conn, error) {
	if strings.HasSuffix(network, "tcp") {
		return h.__dial(ctx, network, addr)
	}
	return h.NetDriver.Dial(ctx, network, addr)
}

func (h *httpDriver) __dial(ctx context.Context, network, addr string) (conn net.Conn, err error) {
	conn, err = h.factory.Dial(ctx, network, h.host)
	if nil != err {
		return
	}
	var body = fmt.Sprintf(h.format, addr, addr)
	if _, err = conn.Write(*(*[]byte)(unsafe.Pointer(&body))); nil != err {
		conn.Close()
		conn = nil
		return
	}
	resp, err := http.ReadResponse(bufio.NewReader(conn), nil)
	resp.Body.Close()
	if err != nil {
		conn.Close()
		conn = nil
		return
	}
	if resp.StatusCode != 200 {
		conn.Close()
		conn = nil
		err = fmt.Errorf("Connect server using proxy error, StatusCode [%d]", resp.StatusCode)
	}
	return
}
