# [Golang lua-vm](https://github.com/vlorc/lua-vm)
[英文版](https://github.com/vlorc/lua-vm/blob/master/README.md)
Golang lua迷你项目

[![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![codebeat badge](https://codebeat.co/badges/c41b426c-4121-4dc8-99c2-f1b60574be64)](https://codebeat.co/projects/github-com-vlorc-lua-vm-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/vlorc/gioc)](https://goreportcard.com/report/github.com/vlorc/lua-vm)
[![GoDoc](https://godoc.org/github.com/vlorc/lua-vm?status.svg)](https://godoc.org/github.com/vlorc/lua-vm)
[![Build Status](https://travis-ci.org/vlorc/lua-vm.svg?branch=master)](https://travis-ci.org/vlorc/lua-vm?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/vlorc/lua-vm/badge.svg?branch=master)](https://coveralls.io/github/vlorc/lua-vm?branch=master)

# 库
+ bit（位操作）
+ buffer（缓冲区）
+ crypto（加密）
+ fs（文件）
+ hash（哈希）
+ io（IO）
+ net（网络）
+ regexp（正则）
+ store（存储）

# 特性
+ http/sock5代理层
+ lua池
+ 文件系统

## 安装
	go get -u github.com/vlorc/lua-vm

## 快速开始

* 创建池
```golang
pool.NewLuaPool()
```

* 预加载
```golang
pool.NewLuaPool().Preload(
		pool.Source(strings.NewReader("print('hello')")),
		pool.Value("tobuffer", base.ToBuffer),
)
```

## 例子

* TCP使用

```golang
import (
    "github.com/vlorc/lua-vm/pool"
    "github.com/vlorc/lua-vm/net/tcp"
    "github.com/vlorc/lua-vm/net/base"
)

func main() {
	p := pool.NewLuaPool().Preload(
		pool.Module("net.tcp", tcp.NewTCPFactory(driver.DirectDriver{})),
		pool.Module("buffer", base.BufferFactory{}),
	)
	if err := p.DoFile("demo/tcp.lua"); nil != err {
		println("error: ", err.Error())
	}
}
```


## Lua 例子
+ [tcp](https://github.com/vlorc/lua-vm/blob/master/demo/tcp.lua)
+ [udp](https://github.com/vlorc/lua-vm/blob/master/demo/udp.lua)
+ [dns](https://github.com/vlorc/lua-vm/blob/master/demo/dns.lua)
+ [http](https://github.com/vlorc/lua-vm/blob/master/demo/http.lua)
+ [time](https://github.com/vlorc/lua-vm/blob/master/demo/time.lua)
+ [crypto](https://github.com/vlorc/lua-vm/blob/master/demo/crypto.lua)
+ [regexp](https://github.com/vlorc/lua-vm/blob/master/demo/regexp.lua)
+ [bit](https://github.com/vlorc/lua-vm/blob/master/demo/bit.lua)
+ [file](https://github.com/vlorc/lua-vm/blob/master/demo/file.lua)