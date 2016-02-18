package alidayu

import ()

type CALLBACK interface {
	// Call when append to queue
	// If return false will get error 'ERR_NOT_PASS_CHECK'
	CALLBACK_Check(interface{}) bool

	// Call before request
	// If return false will do nothing and pass
	CALLBACK_Request(interface{}) bool

	// Call after response
	// bool mean that fail to request alidayu API
	// string is a result from alidayu API
	CALLBACK_Response(interface{}, bool, string)

	CALLBACK_Close()
}
