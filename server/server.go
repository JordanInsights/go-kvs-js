package server

import (
	"go-kvs-js/logs"
	"go-kvs-js/store"
	"net/http"
)

func Init() {
	store.Init()
	logs.Init()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logs.RequestLogger.Println(r.RemoteAddr, r.Method, r.URL.Path)
		routes(w, r)
	})
	http.ListenAndServe(":8000", nil)
}
