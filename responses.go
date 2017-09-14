package ionic

import (
	"encoding/json"
	"net/http"
	"time"
)

// IonResponse represents the response structure expected back from the Ion
// Channel API calls
type IonResponse struct {
	Data json.RawMessage `json:"data"`
	Meta Meta            `json:"meta"`
}

// Meta represents the metadata section of an IonResponse
type Meta struct {
	Copyright  string    `json:"copyright"`
	Authors    []string  `json:"authors"`
	Version    string    `json:"version"`
	LastUpdate time.Time `json:"last_update,omitempty"`
	TotalCount int       `json:"total_count,omitempty"`
	Limit      int       `json:"limit,omitempty"`
	Offset     int       `json:"offset,omitempty"`
}

// IonErrorResponse represents an error response from the Ion Channel API
type IonErrorResponse struct {
	Message string   `json:"message"`
	Fields  []string `json:"fields,omitempty"`
	Code    int      `json:"code"`
}

// NewResponse takes a data object, meta object, and desired status code to
// build a new response.  It returns a message encoded as a byte slice and a
// corresponding status code.  The returned message and status code will reflect
// any errors encountered as part of the encoding process.
func NewResponse(data interface{}, meta Meta, status int) ([]byte, int) {
	b, err := json.Marshal(data)
	if err != nil {
		return NewErrorResponse("failed to marshal data to json", nil, http.StatusInternalServerError)
	}

	// Due to an issue (https://github.com/golang/go/issues/14493) with how
	// RawMessage was handled, versions prior to Go 1.8 will base64 encode the
	// data unless we pass the pointer of the IonResponse struct
	b, err = json.Marshal(&IonResponse{Data: b, Meta: meta})
	if err != nil {
		return NewErrorResponse("failed to marshal response to json", nil, http.StatusInternalServerError)
	}

	return b, status
}

// NewErrorResponse takes a message, related fields, and desired status code to
// build a new error response.  It returns an error message encoded as a byte
// slice and a corresponding status code.  The status code returned will be the
// same as passed into the status parameter unless an error is encountered when
// marshalling the error response into JSON, which will then return an Internal
// Server Error status code.
func NewErrorResponse(message string, fields []string, status int) ([]byte, int) {
	errResp := IonErrorResponse{
		Message: message,
		Fields:  fields,
		Code:    status,
	}

	b, err := json.Marshal(errResp)
	if err != nil {
		b = []byte("failed to marshal error response")
		status = http.StatusInternalServerError
	}

	return b, status
}
