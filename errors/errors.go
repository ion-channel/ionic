package errors

import (
	"fmt"
)

// IonError represents an error from the API with the pertinent information
// accessible
type IonError struct {
	Err            error  `json:"error"`
	ResponseBody   string `json:"response_body"`
	ResponseStatus int    `json:"response_status"`
}

// Errors takes a body, status, format, and any additional arguments to create
// an IonError that includes details from the API
func Errors(body string, status int, format string, a ...interface{}) *IonError {
	return &IonError{fmt.Errorf(format, a...), body, status}
}

func (e IonError) Error() string {
	return fmt.Sprintf("ionic: (%v) %v", e.ResponseStatus, e.Err.Error())
}

// Prepend takes a prefix and puts it on the front of the IonError.
func (e *IonError) Prepend(prefix string) {
	e.Err = fmt.Errorf("%v: %v", prefix, e.Err.Error())
}

// Prepend takes a prefix and an error, creates an IonError if the error is not
// already one, and prepends the prefix onto the existing error.
func Prepend(prefix string, err error) *IonError {
	ierr, ok := err.(*IonError)
	if !ok {
		return Errors("", 0, "%v: %v", prefix, err.Error())
	}

	ierr.Prepend(prefix)
	return ierr
}
