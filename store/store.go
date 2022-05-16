package store

import (
	"fmt"
)

type kvs struct {
	Store map[interface{}]info
}

type listInfo struct {
	Key, Owner string
}

type info struct {
	Key, value interface{}
	Owner      string
}

func Init() {
	go monitorRequests()
}

func (kvs kvs) get(key interface{}) (interface{}, error) {
	keyInfo, hasKey := kvs.Store[key]
	switch hasKey {
	case false:
		return nil, StoreErrors["404"]
	default:
		value := keyInfo.value
		return value, nil
	}
}

func (kvs kvs) put(key interface{}, value interface{}, user string) error {
	existingValue, exists := kvs.Store[key]
	var newInfo = info{
		Key:   key,
		Owner: user,
		value: value,
	}

	if exists && existingValue.Owner != user {
		return StoreErrors["forbidden"]
	}

	kvs.Store[key] = newInfo
	return nil
}

func (kvs kvs) delete(key interface{}, user string) error {
	value, hasKey := kvs.Store[key]

	if hasKey && value.Owner == user {
		delete(kvs.Store, key)
		return nil
	} else if hasKey {
		return StoreErrors["auth"]
	}

	return StoreErrors["404"]
}

func (kvs kvs) list() []listInfo {
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

func (kvs kvs) listKey(key interface{}) (listInfo, error) {
	keyInfo, hasKey := kvs.Store[key]
	infoStruct := listInfo{}

	if hasKey {
		stringifiedKey := fmt.Sprintf("%v", key)

		infoStruct = listInfo{
			Key:   stringifiedKey,
			Owner: keyInfo.Owner,
		}
		return infoStruct, nil
	}

	return infoStruct, StoreErrors["404"]

}
