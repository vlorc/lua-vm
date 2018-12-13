package pool

type LuaModule interface {
	To(val interface{}, name ...string) bool
	Call(name string, args ...interface{})
	Method(name string, val interface{}) bool
}
