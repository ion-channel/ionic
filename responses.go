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
	Copyright  string     `json:"copyright"`
	Authors    []string   `json:"authors"`
	Version    string     `json:"version"`
	LastUpdate *time.Time `json:"last_update,omitempty"`
	TotalCount int        `json:"total_count,omitempty"`
	Limit      int        `json:"limit,omitempty"`
	Offset     int        `json:"offset"`
}

// IonErrorResponse represents an error response from the Ion Channel API
type IonErrorResponse struct {
	Message string   `json:"message"`
	Fields  []string `json:"fields,omitempty"`
	Code    int      `json:"code"`
}

// makeIonResponse constructs an IonResponse object for eventual return
func makeIonResponse(data interface{}, meta Meta) (*IonResponse, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if meta.Copyright == "" {
		meta.Copyright = "Copyright 2018 Selection Pressure LLC www.selectpress.net"
	}
	if meta.Authors == nil || len(meta.Authors) < 1 {
		meta.Authors = append(meta.Authors, "Ion Channel Dev Team")
	}
	if meta.Version == "" {
		meta.Version = "v1"
	}
	return &IonResponse{Data: b, Meta: meta}, nil
}

// NewResponse takes a data object, meta object, and desired status code to
// build a new response.  It returns a message encoded as a byte slice and a
// corresponding status code.  The returned message and status code will reflect
// any errors encountered as part of the encoding process.
func NewResponse(data interface{}, meta Meta, status int) ([]byte, int) {
	ionResponse, err := makeIonResponse(data, meta)
	if err != nil {
		return NewErrorResponse(err.Error(), nil, http.StatusInternalServerError)
	}

	b, err := json.Marshal(ionResponse)
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
