package store

type StoreDriver interface {
	Get(key string) string
	Set(key, value string)
	Delete(key string)
	Exist(key string) bool
	Range(callback func(string, string) bool)
}
