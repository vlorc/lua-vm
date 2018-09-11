package crc

import (
	"hash"
	"hash/crc32"
	"hash/crc64"
	"github.com/vlorc/lua-vm/base"
)

type CRC32Factory struct {}
type CRC64Factory struct {}

func(CRC32Factory)New(poly ...uint32) hash.Hash32{
	if 0 == len(poly) {
		return crc32.NewIEEE()
	}
	return crc32.New(crc32.MakeTable(poly[0]))
}

func(CRC32Factory)Sum32(buf ...base.Buffer) uint32{
	h := crc32.NewIEEE()
	for _,v := range buf {
		h.Write(v)
	}
	return h.Sum32()
}

func(CRC64Factory)New(poly ...uint64) hash.Hash64{
	if 0 == len(poly) {
		crc64.New(crc64.MakeTable(crc64.ISO))
	}
	return crc64.New(crc64.MakeTable(poly[0]))
}

func(CRC64Factory)Sum64(buf ...base.Buffer) uint64{
	h := crc64.New(crc64.MakeTable(crc64.ISO))
	for _,v := range buf {
		h.Write(v)
	}
	return h.Sum64()
}