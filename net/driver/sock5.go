package driver

import (
	"context"
	"golang.org/x/net/proxy"
	"net"
	"net/url"
	"reflect"
	"strings"
	vmnet"github.com/vlorc/lua-vm/net"
)

type Sock5Driver struct {
	vmnet.NetDriver
	dialer Dialer
}

type Dialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
	DialWithConn(ctx context.Context, c net.Conn, network, address string) (net.Addr, error)
	Dial(network, address string) (net.Conn, error)
}

func __newSock5Driver(uri *url.URL,parent vmnet.NetDriver) (vmnet.NetDriver,error){
	var auth *proxy.Auth
	if uri.User != nil {
		password,_ := uri.User.Password()
		auth = &proxy.Auth{
			User: uri.User.Username(),
			Password: password,
		}
	}
	sock5,err := proxy.SOCKS5("tcp", uri.Host,auth,nil)
	if nil != err {
		return nil,err
	}
	reflect.Indirect(reflect.ValueOf(sock5)).FieldByName("ProxyDial").Set(reflect.ValueOf(parent.Dial))

	return &Sock5Driver{
		NetDriver: parent,
		dialer: sock5.(Dialer),
	},nil
}

func NewSock5Driver(rawurl string,parent vmnet.NetDriver) (vmnet.NetDriver,error){
	uri, err := url.Parse(rawurl)
	if nil != err {
		return nil, err
	}
	return __newSock5Driver(uri,parent)
}

func(s *Sock5Driver)Dial(ctx context.Context,network, addr string) (net.Conn, error){
	if strings.HasSuffix(network,"tcp") {
		return s.dialer.DialContext(ctx,network,addr)
	}
	return s.NetDriver.Dial(ctx,network,addr)
}
