package main

import (
	"fmt"
	store "go-kvs-js/store"
	"log"
	"net/http"
)

func main() {
	s := store.Init()

	http.HandleFunc("/", handler) // each request calls a handler
	http.HandleFunc("/ping", ping)

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		get(w, r, s)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

func get(w http.ResponseWriter, r *http.Request, s map[interface{}]interface{}) {
	store.Get(w, r, s)
}
