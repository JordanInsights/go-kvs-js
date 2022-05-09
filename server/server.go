package server

import (
	"go-kvs-js/store"
	"net/http"
)

func Init() {
	s := store.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		routes(w, r, s)
	})

	http.ListenAndServe(":8000", nil)
}
