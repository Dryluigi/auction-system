package errors

import "errors"

type ErrorWithData struct {
	Message string
	Data    interface{}
}

func (err *ErrorWithData) Error() string {
	return err.Message
}

func NewErrorWithData(message string, data interface{}) *ErrorWithData {
	return &ErrorWithData{
		Message: message,
		Data:    data,
	}
}

var ErrEntityNotFound error = errors.New("entity not found")
