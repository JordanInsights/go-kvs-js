package store

import (
	"fmt"
	"net/http"
)

func Init() map[interface{}]interface{} {
	s := make(map[interface{}]interface{})

	s["hello"] = "world"
	// fmt.Println("store: ", s)
	// fmt.Println("store[\"hello\"]: ", s["hello"])
	return s
}

func Get(w http.ResponseWriter, r *http.Request, s map[interface{}]interface{}) {
	key := r.URL.Query().Get("key")
	value := s[key]
	fmt.Fprintf(w, "Value: %q\n", value)
}
