package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"github.com/vlorc/lua-vm/base"
)

type SHA1Factory struct {}
type SHA256Factory struct {}
type SHA512Factory struct {}
type MD5Factory struct {}

func __sum(h hash.Hash,buf ...base.Buffer) base.Buffer{
	for _,v := range buf {
		h.Write(v)
	}
	return base.Buffer(h.Sum(nil))
}

func(SHA1Factory)New() hash.Hash{
	return sha1.New()
}

func(SHA1Factory)Sum(buf ...base.Buffer) base.Buffer{
	return __sum(sha1.New(),buf...)
}

func(MD5Factory)New() hash.Hash{
	return md5.New()
}

func(MD5Factory)Sum(buf ...base.Buffer) base.Buffer{
	return  __sum(md5.New(),buf...)
}

func(SHA256Factory)New() hash.Hash{
	return sha256.New()
}

func(SHA256Factory)Sum(buf ...base.Buffer) base.Buffer{
	return __sum(sha256.New(),buf...)
}

func(SHA512Factory)New() hash.Hash{
	return sha512.New()
}

func(SHA512Factory)Sum(buf ...base.Buffer) base.Buffer{
	return __sum(sha512.New(),buf...)
}