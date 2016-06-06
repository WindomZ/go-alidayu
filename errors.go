package alidayu

import "errors"

var (
	ErrNoAppKey        error = errors.New("goalidayu: Key and Secret must not be null!")
	ErrServiceClosed         = errors.New("goalidayu: Service is closed!")
	ErrServiceOverflow       = errors.New("goalidayu: Service is overflow!")
	ErrPostboxBusy           = errors.New("goalidayu: Postbox is busy")
	ErrCourierBusy           = errors.New("goalidayu: Courier is busy")
)

var (
	ErrMessageNil        error = errors.New("goalidayu: Message must not be null!")
	ErrMessageMethod           = errors.New("goalidayu: Invalid message method")
	ErrMessageFormat           = errors.New("goalidayu: Invalid message format")
	ErrMessageSignMethod       = errors.New("goalidayu: Invalid message sign method")
)

func NewAlidayuResponseError(response string) error {
	return errors.New("goalidayu: " + response)
}
