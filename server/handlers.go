package server

import (
	"encoding/json"
	"fmt"
	"go-kvs-js/store"
	"go-kvs-js/utils"
	"io/ioutil"
	"net/http"
)

func put(w http.ResponseWriter, r *http.Request, kvs store.Kvs, user string) {

	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path, "store")

	if keyPresent {
		value, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "no value")
		}

		stringifiedValue := string(value)
		valueStored := kvs.Put(key, stringifiedValue, user)

		switch valueStored {
		case false:
			w.WriteHeader(http.StatusForbidden)
		default:
			fmt.Fprintf(w, "OK")
		}
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

func get(w http.ResponseWriter, r *http.Request, kvs store.Kvs) {
	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path, "store")

	if keyPresent {
		value, hasKey := kvs.Get(key)

		switch hasKey {

		case false:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 key not found")
			return
		default:
			resString := fmt.Sprint(value)
			fmt.Fprintf(w, resString)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 key not found")
}

func delete(w http.ResponseWriter, r *http.Request, kvs store.Kvs, user string) {
	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path, "store")

	if keyPresent {
		err := kvs.Delete(key, user)
		switch err {
		case store.StoreErrors["404"]:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 key not found")
			return
		case store.StoreErrors["auth"]:
			w.WriteHeader(http.StatusForbidden)
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

func listKey(w http.ResponseWriter, r *http.Request, kvs store.Kvs) {
	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path, "list")
	if keyPresent {
		infoStruct, hasKey := kvs.ListKey(key)
		if hasKey {
			json.NewEncoder(w).Encode(infoStruct)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 key not found")
}
