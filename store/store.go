package store

import (
	"fmt"
	"net/http"
)

func Init() map[interface{}]interface{} {
	s := make(map[interface{}]interface{})
	s["hello"] = "world"
	return s
}

func Get(w http.ResponseWriter, r *http.Request, s map[interface{}]interface{}) {
	key := r.URL.Query().Get("key")
	value, hasKey := s[key]

	switch hasKey {
	case false:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 key not found")
	default:
		fmt.Fprintf(w, "%q", value)
	}
}

func Put(w http.ResponseWriter, r *http.Request, s map[interface{}]interface{}) {

}
