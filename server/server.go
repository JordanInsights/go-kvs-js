package server

import (
	"go-kvs-js/logs"
	"go-kvs-js/store"
	"net/http"
)

func Init() {
	store.Init()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logs.LogRequest(r.RemoteAddr, r.Method, r.URL.Path)
		routes(w, r)
	})
	http.ListenAndServe(":8000", nil)
}
