package store

type operation struct {
	key, authorization, httpMethod, storeMethod string
	value                                       interface{}
	responseChannel                             chan interface{}
	errorChannel                                chan error
}

var requests chan operation = make(chan operation)

func monitorRequests() {
	// protected data
	kvs := kvs{make(map[interface{}]info)}

	for op := range requests {
		go func(op operation) {
			switch op.storeMethod {
			case "Get":
				value, err := kvs.get(op.key)
				op.responseChannel <- value
				op.errorChannel <- err
			case "Put":
				err := kvs.put(op.key, op.value, op.authorization)
				op.responseChannel <- nil
				op.errorChannel <- err
			case "Delete":
				err := kvs.delete(op.key, op.authorization)
				op.responseChannel <- nil
				op.errorChannel <- err
			case "ListKey":
				value, err := kvs.listKey(op.key)
				op.responseChannel <- value
				op.errorChannel <- err
			case "List":
				value := kvs.list()
				op.responseChannel <- value
				op.errorChannel <- nil
			}
		}(op)
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
