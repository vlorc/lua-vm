package rand

import (
	"math/rand"
	"time"
	"github.com/vlorc/lua-vm/base"
)

type RandFactory struct{}

var __rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func (RandFactory) New(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func (RandFactory) Shuffle(n int, swap func(i, j int)){
	__rand.Shuffle(n, swap)
}

func (RandFactory) Read(buf base.Buffer) (int,error) {
	return __rand.Read(buf)
}

func (RandFactory) Perm(n int) []int {
	return __rand.Perm(n)
}

func (RandFactory) Int() int {
	return int(__rand.Int63())
}

func (RandFactory) Int8() int {
	return int(__rand.Int31n(128))
}

func (RandFactory) Int16() int {
	return int(__rand.Int31n(32768))
}

func (RandFactory) Int32() int {
	return int(__rand.Int31())
}

func (RandFactory) Int64() int64 {
	return __rand.Int63()
}

func (RandFactory) Uint8() uint {
	return uint(__rand.Int31n(256))
}

func (RandFactory) Uint32() uint {
	return uint(__rand.Int31n(65536))
}

func (RandFactory) Uint64() uint64 {
	return __rand.Uint64()
}

func (RandFactory) Intn(min, max int64) int64 {
	return __rand.Int63n(max-min) + min
}

func (f RandFactory) Float() float64 {
	return f.Float64()
}

func (RandFactory) Float64() float64 {
	return __rand.Float64()
}

func (RandFactory) Float32() float32 {
	return __rand.Float32()
}

func (f RandFactory) Floatn(min, max float64) float64 {
	return f.Float64n(min, max)
}

func (RandFactory) Float64n(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func (RandFactory) Float32n(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
