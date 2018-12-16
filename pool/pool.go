package pool

import (
	"github.com/yuin/gopher-lua"
	"sync"
)

type LuaPool struct {
	pool sync.Pool
	init []func(*lua.LState) error
}

func __makeNew(p *LuaPool, opt ...lua.Options) func() interface{} {
	return func() interface{} {
		L := lua.NewState(opt...)
		for _, v := range p.init {
			v(L)
		}
		return L
	}
}

func NewLuaPool(opt ...lua.Options) *LuaPool {
	p := &LuaPool{}
	p.pool.New = __makeNew(p, opt...)
	return p
}

func (p *LuaPool) Group(l int) (result []*lua.LState) {
	result = make([]*lua.LState, l)
	for i := range result {
		result[i] = p.pool.Get().(*lua.LState)
	}
	return
}

func (p *LuaPool) Get() *lua.LState {
	return p.pool.Get().(*lua.LState)
}

func (p *LuaPool) Put(L *lua.LState) {
	L.SetTop(0)
	p.pool.Put(L)
}

func (p *LuaPool) Close() error {
	return nil
}

func (p *LuaPool) DoFile(file string) error {
	L := p.Get()
	defer p.Put(L)
	return L.DoFile(file)
}

func (p *LuaPool) DoString(source string) error {
	L := p.Get()
	defer p.Put(L)
	return L.DoString(source)
}

func (p *LuaPool) Preload(load ...func(*lua.LState) error) *LuaPool {
	p.init = append(p.init, load...)
	return p
}

func (p *LuaPool) ModuleString(source string) (LuaModule, error) {
	L := p.Get()
	if err := L.DoString(source); nil != err {
		return nil, err
	}
	return __toModule(L), nil
}

func (p *LuaPool) ModuleFile(file string) (LuaModule, error) {
	L := p.Get()
	if err := L.DoFile(file); nil != err {
		return nil, err
	}
	return __toModule(L), nil
}

func __toModule(state *lua.LState) LuaModule {
	table := state.ToTable(-1)
	return &module{
		state: state,
		table: table,
	}
}
