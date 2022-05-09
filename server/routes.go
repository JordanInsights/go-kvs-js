package server

import (
	"net/http"
	"regexp"
)

var pingPattern = regexp.MustCompile(`ping`)
var getPattern = regexp.MustCompile(`get`)
var putPattern = regexp.MustCompile(`store/?([a-zA-z0-9]+)/?`)

func routes(w http.ResponseWriter, r *http.Request, s map[interface{}]interface{}) {
	path := r.URL.Path
	switch {
	case pingPattern.MatchString(path):
		ping(w, r)
	case putPattern.MatchString(path):
		put(w, r, s)
	case getPattern.MatchString(path):
		get(w, r, s)
	}
}
