package udp

import (
	"github.com/vlorc/lua-vm/base"
	"net"
	"time"
)

func (c *UDPConn) Read(buf base.Buffer) (int, error) {
	if nil == c.remote {
		n, _, e := c.ReadFrom(buf)
		return n, e
	}
	return 0, nil
}

func (c *UDPConn) Write(buf base.Buffer) (int, error) {
	if nil != c.remote {
		return c.WriteTo(buf, c.remote)
	}
	return 0, nil
}

func (c *UDPConn) Close(buf base.Buffer) error {
	return c.conn.Close()
}

func (c *UDPConn) ReadFrom(buf base.Buffer) (n int, a net.Addr, e error) {
	if c.readTimeout > 0 {
		c.conn.SetReadDeadline(time.Now().Add(c.readTimeout))
		n, a, e = c.conn.ReadFrom(buf)
		c.conn.SetReadDeadline(time.Time{})
	} else {
		n, a, e = c.conn.ReadFrom(buf)
	}
	return
}

func (c *UDPConn) WriteTo(buf base.Buffer, addr net.Addr) (n int, e error) {
	if c.writeTimeout > 0 {
		c.conn.SetWriteDeadline(time.Now().Add(c.writeTimeout))
		n, e = c.conn.WriteTo(buf, addr)
		c.conn.SetWriteDeadline(time.Time{})
	} else {
		n, e = c.conn.WriteTo(buf, addr)
	}
	return
}
func (c *UDPConn) SetTimeout(timeout int) {
	c.readTimeout = base.Duration(timeout)
	c.writeTimeout = c.readTimeout
}
func (c *UDPConn) SetReadTimeout(timeout int) {
	c.readTimeout = base.Duration(timeout)
}
func (c *UDPConn) SetWriteTimeout(timeout int) {
	c.writeTimeout = base.Duration(timeout)
}
