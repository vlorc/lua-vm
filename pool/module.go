package pool

import (
	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
	"reflect"
	"sync"
)

type module struct {
	state  *lua.LState
	table  *lua.LTable
	method sync.Map
}

func (m *module) To(val interface{}, name ...string) bool {
	if len(name) > 0 {
		return __toValue(m.state.GetField(m.table, name[0]), val)
	}
	return __toValue(m.table, val)
}

func (m *module) Call(name string, args ...interface{}) {
	m.__callByName(name, 0, nil, args...)
}

func (m *module) Method(name string, val interface{}) bool {
	method := m.__method(name)
	if nil == method {
		return false
	}
	typ := reflect.TypeOf(val).Elem()
	build := func(*lua.LState) []reflect.Value {
		return nil
	}
	if typ.NumOut() > 0 {
		build = func(state *lua.LState) []reflect.Value {
			result := make([]reflect.Value, typ.NumOut())
			top := state.GetTop()
			for i, l := 0, typ.NumOut(); i < l; i++ {
				v := reflect.New(typ.Out(i))
				result[i] = v.Elem()
				if i < top {
					__toValue(state.Get(-(i + 1)), v.Interface())
				}
			}
			return result
		}
	}
	reflect.ValueOf(val).Elem().Set(
		reflect.MakeFunc(
			typ,
			func(args []reflect.Value) (results []reflect.Value) {
				var params []interface{}
				if len(args) > 0 {
					params = make([]interface{}, len(args))
					for i := range args {
						if args[i].IsValid() {
							params[i] = args[i].Interface()
						}
					}
				}
				return m.__call(method, typ.NumOut(), build, params...)
			}),
	)
	return true
}

func (m *module) __method(name string) (method *lua.LFunction) {
	if it, ok := m.method.Load(name); ok {
		method = it.(*lua.LFunction)
	} else if f := m.state.GetField(m.table, name); lua.LTFunction == f.Type() {
		method = f.(*lua.LFunction)
		m.method.Store(name, method)
	}
	return
}

func (m *module) __callByName(name string, nret int, result func(*lua.LState) []reflect.Value, args ...interface{}) []reflect.Value {
	return m.__call(m.__method(name), nret, result, args...)
}

func (m *module) __call(method *lua.LFunction, nret int, result func(*lua.LState) []reflect.Value, args ...interface{}) []reflect.Value {
	thread, cancelFunc := m.state.NewThread()
	defer thread.Close()
	if cancelFunc != nil {
		defer cancelFunc()
	}
	defer thread.SetTop(0)
	thread.Push(method)
	thread.Push(m.table)
	for _, v := range args {
		thread.Push(luar.New(thread, v))
	}
	thread.Call(len(args)+1, nret)
	if nil != result {
		return result(thread)
	}
	return nil
}
