package alidayu

import "errors"

var (
	ERR_NOT_PASS_CHECK   error = errors.New("goalidayu: Not pass when check")
	ERR_NO_APP_KEY             = errors.New("goalidayu: Key and Secret must not be null!")
	ERR_MESSAGE_NIL            = errors.New("goalidayu: Message must not be null!")
	ERR_SERVICE_CLOSED         = errors.New("goalidayu: Service is closed!")
	ERR_SERVICE_OVERFLOW       = errors.New("goalidayu: Service is overflow!")
	ERR_POSTBOX_BUSY           = errors.New("goalidayu: Postbox is busy")
	ERR_COURIER_BUSY           = errors.New("goalidayu: Courier is busy")
)

func NewAlidayuResponseError(response string) error {
	return errors.New("goalidayu: " + response)
}
