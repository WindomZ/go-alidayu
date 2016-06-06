package alidayu

import "errors"

var (
	ERR_NOT_PASS_CHECK   error = errors.New("goalidayu: Not pass when check")
	ERR_NO_APP_KEY       error = errors.New("goalidayu: Key and Secret must not be null!")
	ERR_MESSAGE_NIL      error = errors.New("goalidayu: Message must not be null!")
	ERR_SERVICE_CLOSED   error = errors.New("goalidayu: Service is closed!")
	ERR_SERVICE_OVERFLOW error = errors.New("goalidayu: Service is overflow!")
	ERR_POSTBOX_BUSY     error = errors.New("goalidayu: Postbox is busy")
	ERR_COURIER_BUSY     error = errors.New("goalidayu: Courier is busy")
)
