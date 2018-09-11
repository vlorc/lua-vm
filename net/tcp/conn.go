package tcp

import (
	"time"
	"github.com/vlorc/lua-vm/base"
)

func(c *TCPConn)Read(buf base.Buffer) (l int,e error) {
	if c.readTimeout > 0{
		c.conn.SetReadDeadline(time.Now().Add(c.readTimeout))
		l,e = c.conn.Read(buf)
		c.conn.SetReadDeadline(time.Time{})
	} else {
		l,e = c.conn.Read(buf)
	}
	return l,e
}

func(c *TCPConn)Write(buf base.Buffer) (l int,e error) {
	if c.writeTimeout > 0{
		c.conn.SetWriteDeadline(time.Now().Add(c.writeTimeout))
		l,e = c.conn.Write(buf)
		c.conn.SetWriteDeadline(time.Time{})
	} else {
		l,e = c.conn.Write(buf)
	}
	return l,e
}
func(c *TCPConn)SetTimeout(timeout int) {
	c.readTimeout = base.Duration(timeout)
	c.writeTimeout = c.readTimeout
}
func(c *TCPConn)SetReadTimeout(timeout int) {
	c.readTimeout = base.Duration(timeout)
}
func(c *TCPConn)SetWriteTimeout(timeout int) {
	c.writeTimeout = base.Duration(timeout)
}
func(c*TCPConn)Close() error{
	return c.conn.Close()
}
