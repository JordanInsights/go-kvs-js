package server

import (
	"encoding/json"
	"fmt"
	"go-kvs-js/store"
	"go-kvs-js/utils"
	"net/http"
)

func put(w http.ResponseWriter, r *http.Request, kvs store.Kvs) {

	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path)

	if keyPresent {
		value := r.URL.Query().Get("value")
		valueStored := kvs.Put(key, value)

		switch valueStored {
		case false:
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "forbidden")
		default:
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, "success")
		}
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

func get(w http.ResponseWriter, r *http.Request, kvs store.Kvs) {
	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path)

	if keyPresent {
		value, hasKey := kvs.Get(key)
		switch hasKey {
		case false:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 key not found")
			return
		default:
			fmt.Fprintf(w, "%q", value)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 key not found")
}

func delete(w http.ResponseWriter, r *http.Request, kvs store.Kvs) {
	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path)

	if keyPresent {
		success := kvs.Delete(key)
		switch success {
		case false:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 key not found")
			return
		default:
			fmt.Fprintf(w, "%q", "ok")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 key not found")
}

func list(w http.ResponseWriter, r *http.Request, kvs store.Kvs) {
	list := kvs.List()
	json.NewEncoder(w).Encode(list)
}
