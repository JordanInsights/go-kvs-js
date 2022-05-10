package server

import (
	"go-kvs-js/store"
	"net/http"
	"regexp"
)

var pingPattern = regexp.MustCompile(`ping`)
var putGetDeletePattern = regexp.MustCompile(`store/?([a-zA-z0-9]+)/?`)
var listPattern = regexp.MustCompile(`list`)

func routes(w http.ResponseWriter, r *http.Request, kvs store.Kvs) {
	path := r.URL.Path
	switch {
	case pingPattern.MatchString(path):
		ping(w, r)
	case putGetDeletePattern.MatchString(path) && r.Method == http.MethodPut:
		put(w, r, kvs)
	case putGetDeletePattern.MatchString(path) && r.Method == http.MethodGet:
		get(w, r, kvs)
	case putGetDeletePattern.MatchString(path) && r.Method == http.MethodDelete:
		delete(w, r, kvs)
	case listPattern.MatchString(path) && r.Method == http.MethodGet:
		list(w, r, kvs)
	}
}
