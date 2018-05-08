package penname

import (
	"fmt"
	"net/http"
)

type PenName struct {
	Closed         bool
	Written        []byte
	WrittenHeaders []byte
	returnError    error
}

func New() *PenName {
	return &PenName{}
}

func (p *PenName) Close() error {
	if p.returnError != nil {
		return p.returnError
	}

	p.Closed = true
	return nil
}

// Implements the ResponseWriter interface, returning an empty set of headers
// to meet the interface requirements
func (p *PenName) Header() http.Header {
	return http.Header{}
}

// Convencinece method for reseting state.
func (p *PenName) Reset() {
	p.Closed = false
	p.Written = []byte{}
}

// Sets the error that will be returned when actions are attempted.
func (p *PenName) ReturnError(err error) {
	p.returnError = err
}

// Implements the Writer interface, returning an error if returnError is set.
// The contents of what is written is stored in Written for inspection later.
func (p *PenName) Write(b []byte) (n int, err error) {
	if p.returnError != nil {
		return 0, p.returnError
	}

	p.Written = b
	return len(p.Written), nil
}

// Implements the ResponseWriter interface, capturing headers to the same written buffer
func (p *PenName) WriteHeader(i int) {
	b := []byte(fmt.Sprintf("Header: %v", i))
	p.WrittenHeaders = b
	p.Write(b)
}
