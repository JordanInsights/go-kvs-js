package store

// make a type for the key value store
// make the functions methods on the key value store struct

// The concurrency will be occuring in the store,
// this is where you need to worry about it

type kvp struct {
	key, value interface{}
}

type Kvs struct {
	store map[interface{}]interface{}
}

func Init() Kvs {
	kvs := Kvs{make(map[interface{}]interface{})}
	kvs.Put("hello", "world")
	return kvs
}

func (kvs Kvs) Get(key interface{}) (interface{}, bool) {
	// key := r.URL.Query().Get("key")
	value, hasKey := kvs.store[key]
	return value, hasKey
}

func (kvs Kvs) Put(key interface{}, value interface{}) bool {
	kvs.store[key] = value
	return true
}

func (kvs Kvs) Delete(key interface{}) bool {
	_, hasKey := kvs.store[key]
	if hasKey {
		delete(kvs.store, key)
		return true
	}
	return false
}
