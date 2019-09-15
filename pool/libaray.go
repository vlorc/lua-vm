package pool

import (
	"github.com/yuin/gopher-lua"
)

type luaLib struct {
	libName string
	libFunc lua.LGFunction
}

var luaLibs = []luaLib{
	{lua.LoadLibName, lua.OpenPackage},
	{lua.BaseLibName, lua.OpenBase},
	{lua.TabLibName, lua.OpenTable},
	{lua.StringLibName, lua.OpenString},
	{lua.MathLibName, lua.OpenMath},
	{lua.DebugLibName, lua.OpenDebug},
}

func Library() func(ls *lua.LState) error{
	return func(ls *lua.LState) error {
		for _, lib := range luaLibs {
			ls.Push(ls.NewFunction(lib.libFunc))
			ls.Push(lua.LString(lib.libName))
			ls.Call(1, 0)
		}
		return nil
	}
}

