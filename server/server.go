package server

import (
	"go-kvs-js/store"
	"net/http"
)

func Init() {
	store.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		routes(w, r)
	})

	http.ListenAndServe(":8000", nil)
}
