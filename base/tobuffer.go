package base

import (
	"bytes"
	"encoding"
	"encoding/json"
	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
	"unsafe"
)

func toBuffer(L *luar.LState) int {
	buf := __toBuffer(L.LState)
	L.Push(luar.New(L.LState, buf))
	return 1
}

func __toBuffer(L *lua.LState) Buffer {
	top := L.GetTop()
	if top > 1 {
		return __newBufferN(L)
	}
	if 1 == top {
		return __newBuffer1(L.Get(1))
	}
	return nil
}

func __newBufferN(L *lua.LState) Buffer {
	v := L.Get(1)
	switch v.Type() {
	case lua.LTString:
		r := make([]string, L.GetTop())
		for i := L.GetTop(); i > 0; i-- {
			r[i-1] = string(L.Get(i).(lua.LString))
		}
		return __stringBuffer(r...)
	case lua.LTNumber:
		r := make([]int, L.GetTop())
		for i := L.GetTop(); i > 0; i-- {
			r[i-1] = int(L.Get(i).(lua.LNumber))
		}
		return __numberBuffer(r...)
	case lua.LTUserData:
		r := make([]interface{}, L.GetTop())
		for i := L.GetTop(); i > 0; i-- {
			r[i-1] = int(L.Get(i).(lua.LNumber))
		}
		return __writeBuffer(r...)
	}
	return nil
}

func __newBuffer1(v lua.LValue) Buffer {
	switch v.Type() {
	case lua.LTString:
		return __stringBuffer(string(v.(lua.LString)))
	case lua.LTNumber:
		return __allocBuffer(int(v.(lua.LNumber)), 0)
	case lua.LTUserData:
		return __dataBuffer(v.(*lua.LUserData).Value)
	}
	return nil
}

func __dataBuffer(v interface{}) Buffer {
	switch r := v.(type) {
	case *[]byte:
		if nil != r {
			return Buffer(*r)
		}
	case *Buffer:
		if nil != r {
			return Buffer(*r)
		}
	case *string:
		if nil != r {
			return Buffer(*r)
		}
	case []byte:
		return Buffer(r)
	case encoding.BinaryMarshaler:
		if b, err := r.MarshalBinary(); nil != err && len(b) > 0 {
			return Buffer(b)
		}
	case encoding.TextMarshaler:
		if b, err := r.MarshalText(); nil != err && len(b) > 0 {
			return Buffer(b)
		}
	case json.Marshaler:
		if b, err := r.MarshalJSON(); nil != err && len(b) > 0 {
			return Buffer(b)
		}
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

func __numberBuffer(val ...int) Buffer {
	buf := make(Buffer, len(val))
	for i, v := range val {
		buf[i] = byte(v)
	}
	return buf
}

func __writeBuffer(val ...interface{}) Buffer {
	var buf bytes.Buffer
	for _, v := range val {
		if b := __dataBuffer(v); len(b) > 0 {
			buf.Write(b)
		}
	}
	return Buffer(buf.Bytes())
}
