package tcp

import (
	"net"
	"strconv"
)

func(f *TCPFactory)Listen(addr string,args ...int) (*TCPListener,error) {
	return f.__listen("tcp",addr,args...)
}
func(f *TCPFactory)Listen4(addr string,args ...int) (*TCPListener,error) {
	return f.__listen("tcp4",addr,args...)
}
func(f *TCPFactory)Listen6(addr string,args ...int) (*TCPListener,error) {
	return f.__listen("tcp6",addr,args...)
}
func(f *TCPFactory)Connect(addr string,args ...int) (*TCPConn,error) {
	return f.__connect("tcp",addr,args...)
}
func(f *TCPFactory)Connect4(addr string,args ...int) (*TCPConn,error) {
	return f.__connect("tcp4",addr,args...)
}
func(f *TCPFactory)Connect6(addr string,args ...int) (*TCPConn,error) {
	return f.__connect("tcp6",addr,args...)
}

func(f *TCPFactory)__listen(network,addr string,args ...int) (*TCPListener,error) {
	if len(args) > 0 && args[0] >= 0{
		addr = net.JoinHostPort(addr,strconv.Itoa(args[0]))
	}
	ctx,cancel := f.__context(f.listenTimeout,args)
	defer cancel()
	listen,err := f.driver.Listen(ctx,network,addr)
	if nil != err {
		return nil,err
	}
	return &TCPListener{
		listen: listen,
		readTimeout: f.readTimeout,
		writeTimeout: f.writeTimeout,
	},nil
}

func(f *TCPFactory)__connect(network,addr string,args ...int) (*TCPConn,error) {
	if len(args) > 0 && args[0] >= 0{
		addr = net.JoinHostPort(addr,strconv.Itoa(args[0]))
	}
	ctx,cancel := f.__context(f.listenTimeout,args)
	defer cancel()
	conn,err := f.driver.Dial(ctx,network,addr)
	if nil != err {
		return nil,err
	}
	return &TCPConn{conn: conn},nil
}


