package server

import (
	"fmt"
	"go-kvs-js/store"
	"go-kvs-js/utils"
	"net/http"
)

func put(w http.ResponseWriter, r *http.Request, s map[interface{}]interface{}) {

	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path)

	if keyPresent {
		value := r.URL.Query().Get("value")
		valueStored := store.Put(key, value, s)

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

func get(w http.ResponseWriter, r *http.Request, s map[interface{}]interface{}) {
	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path)

	if keyPresent {
		value, hasKey := store.Get(key, s)
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
