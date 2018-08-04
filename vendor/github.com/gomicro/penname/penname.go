// Package penname provides a mock that implements the Closer & Writer
// interfaces for testing
package penname

import (
	"bytes"
	"fmt"
	"net/http"
)

// PenName contains the state of the mock for performing its designated actions
// and capturing details for assertions
type PenName struct {
	Closed         bool
	written        bytes.Buffer
	writtenHeaders bytes.Buffer
	returnError    error
}

// New returns an initialized PenName for use in tests
func New() *PenName {
	return &PenName{}
}

// Close implements the closer interface, returning an error if returnError is
// set.  Whether or not Close is called is stored in Closed for inspection
// later.
func (p *PenName) Close() error {
	if p.returnError != nil {
		return p.returnError
	}

	p.Closed = true
	return nil
}

// Header implements the ResponseWriter interface, returning an empty set of
// headers to meet the interface requirements
func (p *PenName) Header() http.Header {
	return http.Header{}
}

// Reset is a convencinece method for reseting the state of the mock
func (p *PenName) Reset() {
	p.Closed = false
	p.written.Reset()
	p.writtenHeaders.Reset()
}

// ReturnError sets the error that will be returned when actions are attempted
func (p *PenName) ReturnError(err error) {
	p.returnError = err
}

// Write implements the Writer interface, returning an error if returnError is
// set.  The contents of what is written is stored in Written for inspection
// later.
func (p *PenName) Write(b []byte) (n int, err error) {
	if p.returnError != nil {
		return 0, p.returnError
	}

	p.written.Write(b)
	return len(b), nil
}

// Written returns the bytes that have been written to the writer
func (p *PenName) Written() []byte {
	return p.written.Bytes()
}

// WriteHeader implements the ResponseWriter interface, capturing headers to a
// separate
func (p *PenName) WriteHeader(i int) {
	p.writtenHeaders.WriteString(fmt.Sprintf("Header: %v", i))
}

// WrittenHeaders returns the bytes of the headers that have been written to
// the writer
func (p *PenName) WrittenHeaders() []byte {
	return p.writtenHeaders.Bytes()
}
