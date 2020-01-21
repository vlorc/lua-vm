package base

import (
	"bytes"
	luar "layeh.com/gopher-luar"
	"unicode/utf8"
	"unsafe"
)

type Buffer []byte

type BufferFactory struct{}

func (f BufferFactory) New(length int) Buffer {
	return make(Buffer, length)
}

func (BufferFactory) Alloc(length int, args ...int) Buffer {
	return __allocBuffer(length, args...)
}

func (BufferFactory) Form(L *luar.LState) int {
	return toBuffer(L.LState)
}

func (BufferFactory) FormNumber(val ...int) Buffer {
	return __numberBuffer(val...)
}

func (BufferFactory) FormString(val ...string) Buffer {
	return __stringBuffer(val...)
}

func (b Buffer) Split(L *luar.LState) int {
	sep := __toBuffer(L.LState)
	ret := bytes.Split(b, sep)

	L.Push(luar.New(L.LState, *(*[]Buffer)(unsafe.Pointer(&ret))))
	return 1
}

func (b Buffer) Slice(args ...int) Buffer {
	begin, end := 0, len(b)
	if len(args) > 0 {
		if begin = args[0] - 1; len(args) > 1 {
			end = args[1]
		}
	}
	return b[begin:end]
}

func (b Buffer) Peek(i int) Buffer {
	if i > 0 {
		if i >= len(b) {
			return Buffer{}
		}
		return b[i-1:]
	}
	return b
}

func (b Buffer) IndexAny(str string) int {
	return bytes.IndexAny(b, str) + 1
}

func (b Buffer) IndexByte(val int) int {
	return bytes.IndexByte(b, byte(val)) + 1
}

func (b Buffer) Index(L *luar.LState) int {
	sep := __newBufferN(L.LState)
	ret := bytes.Index(b, sep) + 1

	L.Push(luar.New(L.LState, ret))
	return 1
}

func (b Buffer) IndexString(str string) int {
	return bytes.Index(b, *(*[]byte)(unsafe.Pointer(&str))) + 1
}

func (b Buffer) LastAny(str string) int {
	return bytes.LastIndexAny(b, str) + 1
}

func (b Buffer) LastByte(val int) int {
	return bytes.LastIndexByte(b, byte(val)) + 1
}

func (b Buffer) Last(L *luar.LState) int {
	sep := __newBufferN(L.LState)
	ret := bytes.LastIndex(b, sep) + 1

	L.Push(luar.New(L.LState, ret))
	return 1
}

func (b Buffer) LastString(str string) int {
	return bytes.LastIndex(b, *(*[]byte)(unsafe.Pointer(&str))) + 1
}

func (b Buffer) Clone(args ...int) Buffer {
	src := b.Slice(args...)
	dst := make(Buffer, len(src))
	copy(dst, src)
	return dst
}

func (b Buffer) Copy(src Buffer, args ...int) int {
	return copy(b.Slice(args...), src)
}

func (b Buffer) Concat(src ...Buffer) Buffer {
	l := len(b)
	for _, v := range src {
		l += len(v)
	}
	dst := make(Buffer, l)
	l = copy(dst, b)
	for _, v := range src {
		l += copy(dst[l:], v)
	}
	return dst
}

func (b Buffer) Equal(buf Buffer) bool {
	return bytes.Equal(b, buf)
}

func (b Buffer) String() string {
	return __HEXString(b)
}

func (b Buffer) ToString(args ...string) string {
	if len(args) <= 0 {
		return __rawString(b)
	}
	if encode, ok := EncodeTable[args[0]]; ok {
		return encode(b)
	}
	return ""
}

func (b Buffer) Reverse(args ...int) Buffer {
	n := b.Slice(args...)
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	return n
}

func (b Buffer) ToChar(args ...int) int {
	n := b.Slice(args...)
	r, _ := utf8.DecodeRune(n)
	return int(r)
}

func (b Buffer) ToNumber(args ...int) uint64 {
	n := b.Slice(args...)
	i := 0
	if len(args) > 2 && args[2] >= 0 && args[2] < len(__number) {
		i = args[2]
	}
	return __number[i](n)
}

func (b Buffer) ToLine(args ...int) string {
	n := b.Slice(args...)
	if len(n) <= 0 {
		return ""
	}
	pos := bytes.IndexByte(n, byte('\n'))
	if pos > 0 {
		n = n[:pos]
	}
	return __rawString(n)
}

func (b Buffer) ToHash(args ...string) uint64 {
	if len(args) <= 0 {
		return __hashChecksum8(b)
	}
	if hash, ok := HashTable[args[0]]; ok {
		return hash(b)
	}
	return 0
}
