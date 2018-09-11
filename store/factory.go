package store

type StoreFactory struct {
	factory map[string]func(...string) StoreDriver
}

var Default = map[string]func(...string) StoreDriver{
	"memory": func(...string) StoreDriver {
		return &MemoryStore{}
	},
}

func NewStoreFactory(factory map[string]func(...string) StoreDriver) *StoreFactory {
	if nil == factory {
		factory = Default
	} else {
		for k, v := range Default {
			if _, ok := factory[k]; !ok {
				factory[k] = v
			}
		}
	}
	return &StoreFactory{
		factory: factory,
	}
}

func (f *StoreFactory) New(name string, args ...string) StoreDriver {
	factory, ok := f.factory[name]
	if !ok {
		return nil
	}
	return factory(args...)
}
