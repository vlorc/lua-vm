package store

import "sync"

type MemoryStore struct {
	m sync.Map
}

type MemoryStoreFactory struct{}

func(MemoryStoreFactory)New() StoreDriver{
	return &MemoryStore{}
}

func(s *MemoryStore)Get(key string) string{
	tmp,ok := s.m.Load(key)
	if !ok {
		return ""
	}
	return tmp.(string)
}
func(s *MemoryStore)Set(key,value string) {
	s.m.Store(key,value)
}
func(s *MemoryStore)Delete(key string) {
	s.m.Delete(key)
}
func(s *MemoryStore)Exist(key string) bool {
	_,ok := s.m.Load(key)
	return ok
}
func(s *MemoryStore)Range(callback func(string,string)bool) {
	s.m.Range(func(key, value interface{}) bool {
		return callback(key.(string),value.(string))
	})
}