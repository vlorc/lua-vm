package store

type StoreDriver interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Delete(key string)
	Exist(key string) bool
	Range(callback func(string, interface{}) bool)
}
