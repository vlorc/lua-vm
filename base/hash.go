package base

import (
	"hash/crc32"
	"hash/crc64"
)

var HashTable = map[string]func(Buffer) uint64{
	"checksum8":  __hashChecksum8,
	"checksum16": __hashChecksum16,
	"checksum32": __hashChecksum32,
	"checksum64": __hashChecksum64,
	"crc32":      __hashCrc32,
	"crc64":      __hashCrc64ISO,
	"crc64.iso":  __hashCrc64ISO,
	"crc64.ecma": __hashCrc64ECMA,
}

func __hashChecksum8(b Buffer) uint64 {
	r := byte(0)
	for _, v := range b {
		r += v
	}
	return uint64(r)
}

func __hashChecksum16(b Buffer) uint64 {
	r := uint16(0)
	for _, v := range b {
		r += uint16(v)
	}
	return uint64(r)
}

func __hashChecksum32(b Buffer) uint64 {
	r := uint32(0)
	for _, v := range b {
		r += uint32(v)
	}
	return uint64(r)
}

func __hashChecksum64(b Buffer) uint64 {
	r := uint64(0)
	for _, v := range b {
		r += uint64(v)
	}
	return uint64(r)
}

func __hashCrc32(b Buffer) uint64 {
	h := crc32.NewIEEE()
	h.Write(b)
	return uint64(h.Sum32())
}

func __hashCrc64ISO(b Buffer) uint64 {
	h := crc64.New(crc64.MakeTable(crc64.ISO))
	h.Write(b)
	return h.Sum64()
}

func __hashCrc64ECMA(b Buffer) uint64 {
	h := crc64.New(crc64.MakeTable(crc64.ECMA))
	h.Write(b)
	return h.Sum64()
}
