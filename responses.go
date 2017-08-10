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
