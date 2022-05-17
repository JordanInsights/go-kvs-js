package server

import (
	"encoding/json"
	"fmt"
	"go-kvs-js/store"
	"go-kvs-js/utils"
	"io/ioutil"
	"net/http"
	"os"
)

// func put(w http.ResponseWriter, r *http.Request, kvs store.Kvs, user string) {
func put(w http.ResponseWriter, r *http.Request, user string) {
	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path, "store")

	if keyPresent {
		value, valueErr := ioutil.ReadAll(r.Body)

		if valueErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "no value")
			return
		}

		stringifiedValue := string(value)

		_, err := store.AddRequest(user, key, r.Method, "Put", stringifiedValue)

		switch err {
		case store.StoreErrors["forbidden"]:
			w.WriteHeader(http.StatusForbidden)
		case nil:
			fmt.Fprintf(w, "ok")
		}
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

// func get(w http.ResponseWriter, r *http.Request, kvs store.Kvs, user string) {
func get(w http.ResponseWriter, r *http.Request, user string) {
	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path, "store")

	if keyPresent {
		res, err := store.AddRequest(user, key, r.Method, "Get", nil)

		switch err {
		case nil:
			resString := fmt.Sprint(res)
			fmt.Fprint(w, resString)
			return
		case store.StoreErrors["404"]:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 key not found")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 key not found")
}

// func delete(w http.ResponseWriter, r *http.Request, kvs store.Kvs, user string) {
func delete(w http.ResponseWriter, r *http.Request, user string) {
	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path, "store")

	if keyPresent {
		_, err := store.AddRequest(user, key, r.Method, "Delete", nil)

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

// func list(w http.ResponseWriter, r *http.Request, kvs store.Kvs) {
func list(w http.ResponseWriter, r *http.Request) {
	res, _ := store.AddRequest("", "", r.Method, "List", nil)
	json.NewEncoder(w).Encode(res)
}

// func listKey(w http.ResponseWriter, r *http.Request, kvs store.Kvs) {
func listKey(w http.ResponseWriter, r *http.Request) {
	key, keyPresent := utils.ExtractParametricEndpoints(r.URL.Path, "list")

	if keyPresent {
		res, err := store.AddRequest("", key, r.Method, "ListKey", nil)

		switch err {
		case store.StoreErrors["404"]:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 key not found")
		default:
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 key not found")
}

func shutdown(w http.ResponseWriter, r *http.Request, u string) {
	if u != "admin" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
	go func() {
		store.StopRequests()
		os.Exit(0)
	}()
}
