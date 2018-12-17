package base

import (
	"math/bits"
)

type BitFactory struct{}

func (BitFactory) Size() int {
	return bits.UintSize
}

func (BitFactory) Left(v, r uint) uint {
	return v << r
}

func (BitFactory) Right(v, r uint) uint {
	return v >> r
}

func (BitFactory) Right8(v, r uint8) uint8 {
	return v >> r
}

func (BitFactory) Left8(v, r uint8) uint8 {
	return v << r
}

func (BitFactory) Right16(v, r uint16) uint16 {
	return v >> r
}

func (BitFactory) Left16(v, r uint16) uint16 {
	return v << r
}

func (BitFactory) Right32(v, r uint) uint {
	return v >> r
}

func (BitFactory) Left32(v, r uint32) uint32 {
	return v << r
}

func (BitFactory) Left64(v, r uint64) uint64 {
	return v << r
}

func (BitFactory) Right64(v, r uint64) uint64 {
	return v >> r
}

func (BitFactory) Xor(v, r uint) uint {
	return v ^ r
}

func (BitFactory) And(v, r uint) uint {
	return v & r
}

func (BitFactory) Or(v, r uint) uint {
	return v | r
}

func (BitFactory) Not(v uint) uint {
	return ^v
}

func (BitFactory) Test(v, r uint) bool {
	return 0 != (v>>r)&1
}

func (BitFactory) Count(v uint) int {
	return bits.OnesCount(v)
}

func (BitFactory) RotateLeft(v uint, k int) uint {
	return bits.RotateLeft(v, k)
}

func (BitFactory) RotateLeft8(v uint8, k int) uint8 {
	return bits.RotateLeft8(v, k)
}

func (BitFactory) RotateLeft16(v uint16, k int) uint16 {
	return bits.RotateLeft16(v, k)
}

func (BitFactory) RotateLeft32(v uint32, k int) uint32 {
	return bits.RotateLeft32(v, k)
}

func (BitFactory) RotateLeft64(v uint64, k int) uint64 {
	return bits.RotateLeft64(v, k)
}

func (BitFactory) Reverse(v uint) uint {
	return bits.Reverse(v)
}

func (BitFactory) Reverse8(v uint8) uint8 {
	return bits.Reverse8(v)
}

func (BitFactory) Reverse16(v uint16) uint16 {
	return bits.Reverse16(v)
}

func (BitFactory) Reverse32(v uint32) uint32 {
	return bits.Reverse32(v)
}

func (BitFactory) Reverse64(v uint64) uint64 {
	return bits.Reverse64(v)
}
