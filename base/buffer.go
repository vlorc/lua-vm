package base

import (
	"bytes"
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

func (BufferFactory) Form(str string) Buffer {
	return Buffer(str)
}

func (BufferFactory) FormNumber(val ...int) Buffer {
	return __numberBuffer(val...)
}

func (BufferFactory) FormString(val ...string) Buffer {
	return __stringBuffer(val...)
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

func (b Buffer) IndexAny(str string) int {
	return bytes.IndexAny(b, str) + 1
}

func (b Buffer) IndexByte(val int) int {
	return bytes.IndexByte(b, byte(val)) + 1
}

func (b Buffer) Index(buf Buffer) int {
	return bytes.Index(b, buf) + 1
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

func (b Buffer) Last(buf Buffer) int {
	return bytes.LastIndex(b, buf) + 1
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
	i := len(b)
	for _, v := range src {
		i += len(v)
	}
	dst := make(Buffer, i)
	i = copy(dst, b)
	for _, v := range src {
		i += copy(dst[i:], v)
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

func (b Buffer) ToRune(args ...int) rune {
	n := b.Slice(args...)
	r, _ := utf8.DecodeRune(n)
	return r
}

func (b Buffer) ToNumber(args ...int) (r uint64) {
	n := b.Slice(args...)
	if 3 != len(args) {
		for i := len(n) - 1; i >= 0; i-- {
			r = (r << 8) + uint64(n[i])
		}
	} else {
		for _, v := range n {
			r = (r << 8) + uint64(v)
		}
	}
	return r
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
