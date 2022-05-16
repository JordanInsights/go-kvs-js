package server

import (
	"go-kvs-js/store"
	"net/http"
)

func Init() {
	// kvs := store.Init()
	store.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// routes(w, r, kvs)
		routes(w, r)
	})

	http.ListenAndServe(":8000", nil)
}
