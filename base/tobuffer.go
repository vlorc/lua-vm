package base

import (
	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
	"reflect"
	"unsafe"
)

func ToBuffer(L *lua.LState) int {
	buf := __toBuffer(L)
	L.Push(luar.New(L, buf))
	return 1
}

func __toBuffer(L *lua.LState) Buffer {
	top := L.GetTop()
	if 1 == top {
		return __newBuffer(L.Get(1))
	}
	return nil
}

func __newBuffer(v lua.LValue) Buffer {
	switch v.Type() {
	case lua.LTString:
		__stringBuffer(string(v.(lua.LString)))
	case lua.LTNumber:
		__allocBuffer(int(v.(lua.LNumber)), 0)
	case lua.LTUserData:
		__dataBuffer(v.(*lua.LUserData).Value)
	}
	return nil
}

func __dataBuffer(v interface{}) Buffer {
	reflect.TypeOf(v)
	switch r := v.(type) {
	case []byte:
		return Buffer(r)
	}
	return nil
}

func __allocBuffer(length int, args ...int) Buffer {
	b := make(Buffer, length)
	if len(args) > 0 {
		b[0] = byte(args[0])
		for i := 1; i < len(b); i *= 2 {
			copy(b[i:], b[:i])
		}
	}
	return b
}

func __stringBuffer(val ...string) Buffer {
	i := 0
	for i := range val {
		i += len(val[i])
	}
	buf := make(Buffer, i)
	i = 0
	for i, v := range val {
		if len(v) > 0 {
			copy(buf[i:], *(*[]byte)(unsafe.Pointer(&v)))
			i += len(v)
		}
	}
	return buf
}

func __intBuffer(val ...int) Buffer {
	buf := make(Buffer, len(val))
	for i, v := range val {
		buf[i] = byte(v)
	}
	return buf
}
