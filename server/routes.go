package server

import (
	"net/http"
	"regexp"
)

var pingPattern = regexp.MustCompile(`ping`)
var putGetPattern = regexp.MustCompile(`store/?([a-zA-z0-9]+)/?`)

func routes(w http.ResponseWriter, r *http.Request, s map[interface{}]interface{}) {
	path := r.URL.Path
	switch {
	case pingPattern.MatchString(path):
		ping(w, r)
	case putGetPattern.MatchString(path) && r.Method == http.MethodPut:
		put(w, r, s)
	case putGetPattern.MatchString(path) && r.Method == http.MethodGet:
		get(w, r, s)
	}
}
