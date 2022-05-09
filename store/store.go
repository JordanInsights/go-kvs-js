package store

func Init() map[interface{}]interface{} {
	s := make(map[interface{}]interface{})
	s["hello"] = "world"
	return s
}

func Get(key interface{}, s map[interface{}]interface{}) (interface{}, bool) {
	// key := r.URL.Query().Get("key")
	value, hasKey := s[key]
	return value, hasKey
}

func Put(key interface{}, value interface{}, s map[interface{}]interface{}) bool {
	s[key] = value
	return true
}
