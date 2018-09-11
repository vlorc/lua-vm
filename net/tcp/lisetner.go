package tcp

import "github.com/vlorc/lua-vm/base"

func(l *TCPListener)Accept() (*TCPConn,error) {
	conn,err := l.listen.Accept()
	if nil != err {
		return nil,err
	}
	return &TCPConn{
		conn: conn,
		readTimeout: l.readTimeout,
		writeTimeout: l.writeTimeout,
	},nil
}
func(l *TCPListener)Close() error {
	return l.listen.Close()
}
func(l *TCPListener)SetTimeout(timeout int) {
	l.readTimeout = base.Duration(timeout)
	l.writeTimeout = l.readTimeout
}
func(l *TCPListener)SetReadTimeout(timeout int) {
	l.readTimeout = base.Duration(timeout)
}
func(l *TCPListener)SetWriteTimeout(timeout int) {
	l.writeTimeout = base.Duration(timeout)
}
