package store

import (
	"encoding/json"
	"errors"
	"fmt"
)

// make a type for the key value store
// make the functions methods on the key value store struct

// The concurrency will be occuring in the store,
// this is where you need to worry about it

type kvp struct {
	key, value interface{}
}

type Kvs struct {
	Store map[interface{}]info
}

type listInfo struct {
	Key, Owner string
}

type info struct {
	Key, value interface{}
	Owner      string
}

func Init() Kvs {
	kvs := Kvs{make(map[interface{}]info)}
	return kvs
}

func (kvs Kvs) Get(key interface{}) (interface{}, bool) {
	keyInfo, hasKey := kvs.Store[key]
	value := keyInfo.value
	return value, hasKey
}

func (kvs Kvs) Put(key interface{}, value interface{}) bool {
	var newInfo = info{
		Key:   key,
		Owner: "Jordan",
		value: value,
	}
	kvs.Store[key] = newInfo
	return true
}

func (kvs Kvs) Delete(key interface{}) bool {
	_, hasKey := kvs.Store[key]
	if hasKey {
		delete(kvs.Store, key)
		return true
	}
	return false
}

func (kvs Kvs) List() []listInfo {

	var convertedStore []listInfo
	for key, info := range kvs.Store {
		stringifiedKey := fmt.Sprintf("%v", key)

		infoStruct := listInfo{
			Key:   stringifiedKey,
			Owner: info.Owner,
		}

		convertedStore = append(convertedStore, infoStruct)
	}

	return convertedStore
}

func (kvs Kvs) ListKey(key interface{}) ([]byte, error) {
	keyInfo, hasKey := kvs.Store[key]
	if hasKey {
		type keyAndOwner struct {
			Key   interface{}
			Owner string
		}

		returnVal := keyAndOwner{
			Key:   key,
			Owner: keyInfo.Owner,
		}

		json, err := json.Marshal(returnVal)
		return json, err
	}

	placeholder, _ := json.Marshal("Error")
	return placeholder, errors.New("Error")
}
