package store

import "errors"

var StoreErrors = map[string]error{
	"auth":      errors.New("Unauthorized"),
	"404":       errors.New("404 Item not found"),
	"forbidden": errors.New("forbidden"),
}
