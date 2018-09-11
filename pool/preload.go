package pool

import (
	"github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"io"
	"layeh.com/gopher-luar"
)

func Source(r io.Reader, name ...string) func(*lua.LState) error {
	if len(name) > 0 && len(name[0]) > 0 {
		return source(r, name[0])
	}
	return source(r, "<string>")
}

func source(r io.Reader, name string) func(*lua.LState) error {
	chunk, err := parse.Parse(r, name)
	if nil != err {
		return func(*lua.LState) error {
			return err
		}
	}
	proto, err := lua.Compile(chunk, name)
	if err != nil {
		return func(*lua.LState) error {
			return err
		}
	}
	return func(L *lua.LState) error {
		L.Push(L.NewFunctionFromProto(proto))
		return L.PCall(0, lua.MultRet, nil)
	}
}

func Module(name string, val interface{}) func(*lua.LState) error {
	return func(L *lua.LState) error {
		L.PreloadModule(name, func(S *lua.LState) int {
			S.Push(luar.New(S, val))
			return 1
		})
		return nil
	}
}

func Type(name string, val interface{}) func(*lua.LState) error {
	return func(L *lua.LState) error {
		L.SetGlobal(name, luar.NewType(L, val))
		return nil
	}
}

func Value(name string, val interface{}) func(*lua.LState) error {
	return func(L *lua.LState) error {
		L.SetGlobal(name, luar.New(L, val))
		return nil
	}
}
