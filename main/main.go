package main

import (
	"github.com/vlorc/lua-vm/base"
	"github.com/vlorc/lua-vm/crypto/hash"
	"github.com/vlorc/lua-vm/crypto/rand"
	"github.com/vlorc/lua-vm/fs"
	"github.com/vlorc/lua-vm/hash/crc"
	"github.com/vlorc/lua-vm/io"
	"github.com/vlorc/lua-vm/net/dns"
	"github.com/vlorc/lua-vm/net/driver"
	"github.com/vlorc/lua-vm/net/http"
	"github.com/vlorc/lua-vm/net/tcp"
	"github.com/vlorc/lua-vm/net/udp"
	"github.com/vlorc/lua-vm/net/url"
	"github.com/vlorc/lua-vm/pool"
	"github.com/vlorc/lua-vm/regexp"
	"github.com/vlorc/lua-vm/store"
	"time"
)

func main() {
	network := driver.DirectDriver{}
	p := pool.NewLuaPool().Preload(
		pool.Value("tobuffer", base.ToBuffer),
		pool.Module("net.tcp", tcp.NewTCPFactory(network)),
		pool.Module("net.udp", udp.NewUDPFactory(network)),
		pool.Module("net.http", http.NewHTTPFactory(network)),
		pool.Module("net.dns", dns.NewDNSFactory(network)),
		pool.Module("buffer", base.BufferFactory{}),
		pool.Module("time", base.TimeFactory{}),
		pool.Module("bit", base.BitFactory{}),
		pool.Module("fs", fs.NewRelativeFileFactory(".",fs.NativeFileFactory{})),
		pool.Module("io.reader", io.ReaderFactory{}),
		pool.Module("io.writer", io.WriterFactory{}),
		pool.Module("net.url", url.URlFactory{}),
		pool.Module("crypto.rand", rand.RandFactory{}),
		pool.Module("crypto.md5", hash.MD5Factory{}),
		pool.Module("crypto.sha1", hash.SHA1Factory{}),
		pool.Module("crypto.sha256", hash.SHA256Factory{}),
		pool.Module("crypto.sha512", hash.SHA512Factory{}),
		pool.Module("hash.crc32", crc.CRC32Factory{}),
		pool.Module("hash.crc64", crc.CRC64Factory{}),
		pool.Module("regexp", regexp.RegexpFactory{}),
		pool.Module("store",store.NewStoreFactory(nil)),
	)

	now := time.Now()
	err := p.DoFile("demo/tcp.lua")
	if nil != err {
		println("error: ", err.Error())
	}
	last := time.Now()
	println("time: ", (last.UnixNano()-now.UnixNano())/1000000, last.Unix()-now.Unix())
}
