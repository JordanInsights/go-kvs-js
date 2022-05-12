package server

import (
	"go-kvs-js/store"
	"net/http"
	"regexp"
)

var pingPattern = regexp.MustCompile(`ping`)
var putGetDeletePattern = regexp.MustCompile(`store/?([a-zA-z0-9]+)/?`)
var listKeyPattern = regexp.MustCompile(`list/{1}([a-zA-z0-9]+)/?`)
var listPattern = regexp.MustCompile(`list`)
var shutdownPattern = regexp.MustCompile(`shutdown`)

func routes(w http.ResponseWriter, r *http.Request, kvs store.Kvs) {
	path := r.URL.Path
	user := r.Header.Get("Authorization")

	switch {
	case pingPattern.MatchString(path):
		ping(w, r)
	case putGetDeletePattern.MatchString(path) && r.Method == http.MethodPut:
		put(w, r, kvs, user)
	case putGetDeletePattern.MatchString(path) && r.Method == http.MethodGet:
		get(w, r, kvs)
	case putGetDeletePattern.MatchString(path) && r.Method == http.MethodDelete:
		delete(w, r, kvs, user)
	case listKeyPattern.MatchString(path) && r.Method == http.MethodGet:
		listKey(w, r, kvs)
	case listPattern.MatchString(path) && r.Method == http.MethodGet:
		list(w, r, kvs)
	case shutdownPattern.MatchString(path) && r.Method == http.MethodGet:
		shutdown(w, r, user)
	}
}
