package store

// type operation struct {
// 	action     string
// 	key, value interface{}
// }

type operation struct {
	key, authorization, httpMethod, storeMethod string
	value                                       interface{}
	responseChannel                             chan interface{}
	errorChannel                                chan error
}

var requests chan operation = make(chan operation)

func monitorRequests() {
	// protected data
	kvs := Kvs{make(map[interface{}]info)}

	for op := range requests {
		switch op.storeMethod {
		case "Get":
			value, err := kvs.Get(op.key)
			op.responseChannel <- value
			op.errorChannel <- err
		case "Put":
			err := kvs.Put(op.key, op.value, op.authorization)
			op.responseChannel <- nil
			op.errorChannel <- err
		case "Delete":
			err := kvs.Delete(op.key, op.authorization)
			op.responseChannel <- nil
			op.errorChannel <- err
		case "ListKey":
			value, err := kvs.ListKey(op.key)
			op.responseChannel <- value
			op.errorChannel <- err
		case "List":
			value := kvs.List()
			op.responseChannel <- value
			op.errorChannel <- nil
		}
	}
}

func AddRequest(authorization string, key string, httpMethod string, storeMethod string, value interface{}) (interface{}, error) {
	errorChannel := make(chan error)
	responseChannel := make(chan interface{})

	op := operation{
		authorization:   authorization,
		key:             key,
		httpMethod:      httpMethod,
		storeMethod:     storeMethod,
		value:           value,
		errorChannel:    errorChannel,
		responseChannel: responseChannel,
	}

	requests <- op

	return <-responseChannel, <-errorChannel
}
