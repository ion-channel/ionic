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

// Errorf takes a body, status, format, and any additional arguments to create
// an IonError that includes details from the API
func Errorf(body string, status int, format string, a ...interface{}) *IonError {
	return &IonError{fmt.Errorf(format, a...), body, status}
}

func (e IonError) Error() string {
	return fmt.Sprintf("ionic: (%v) %v", e.ResponseStatus, e.Err)
}
