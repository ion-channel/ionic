package ionic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// IonResponse represents the response structure expected back from the Ion
// Channel API calls
type IonResponse struct {
	Data   json.RawMessage `json:"data"`
	Meta   Meta            `json:"meta"`
	status int
}

// Meta represents the metadata section of an IonResponse
type Meta struct {
	Copyright  string     `json:"copyright"`
	Authors    []string   `json:"authors"`
	Version    string     `json:"version"`
	LastUpdate *time.Time `json:"last_update,omitempty"`
	TotalCount int        `json:"total_count"`
	Limit      int        `json:"limit,omitempty"`
	Offset     int        `json:"offset"`
}

// NewResponse takes a data object, meta object, and desired status code to
// build a new response.  It returns a newly formed Ion Response object with
// default values applied to the response.
func NewResponse(data interface{}, meta Meta, status int) (*IonResponse, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %v", err.Error())
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

	return &IonResponse{Data: b, Meta: meta, status: status}, nil
}

// WriteResponse takes an http ResponseWriter and writes the Ion Response object
// to the writer. It encodes the object as marshalled JSON. If it encounters any
// errors while attempt to formulate the response, it will write an Ion Error
// Response to the writer instead.
func (ir *IonResponse) WriteResponse(w http.ResponseWriter) {
	b, err := json.Marshal(ir)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal response to json: %v", err.Error())
		er := NewErrorResponse(errMsg, nil, http.StatusInternalServerError)
		er.WriteResponse(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ir.status)
	fmt.Fprintln(w, string(b))
}

// IonErrorResponse represents an error response from the Ion Channel API
type IonErrorResponse struct {
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
	Code    int               `json:"code"`
}

// NewErrorResponse takes a message, related fields, and desired status code to
// build a new error response.  It returns an error message encoded as a byte
// slice and a corresponding status code.  The status code returned will be the
// same as passed into the status parameter unless an error is encountered when
// marshalling the error response into JSON, which will then return an Internal
// Server Error status code.
func NewErrorResponse(message string, fields map[string]string, status int) *IonErrorResponse {
	return &IonErrorResponse{
		Message: message,
		Fields:  fields,
		Code:    status,
	}
}

// WriteResponse takes an http ResponseWriter and writes the Ion Error Response
// object to the writer. It encodes the object as marshalled JSON. If it
// encounters any errors while attempt to formulate the response, it will write
// a standard message and internal server error status to the writer.
func (er *IonErrorResponse) WriteResponse(w http.ResponseWriter) {
	b, err := json.Marshal(er)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "failed to marshal error response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(er.Code)
	fmt.Fprintln(w, string(b))
}
