// Package ionic provides a direct representation of the endpoints and objects
// within the Ion Channel API
package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/requests"
)

const (
	maxIdleConns        = 25
	maxIdleConnsPerHost = 25
	maxPagingLimit      = 100
)

// IonClient represnets a communication layer with the Ion Channel API
type IonClient struct {
	baseURL *url.URL
	client  *http.Client
}

// New takes the base URL of the API and returns a client for talking to the API
// and an error if any issues instantiating the client are encountered
func New(baseURL string) (*IonClient, error) {
	c := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnsPerHost,
			MaxIdleConns:        maxIdleConns,
		},
	}

	return NewWithClient(baseURL, c)
}

// NewWithClient takes the base URL of the API and an existing HTTP client.  It
// returns a client for talking to the API and an error if any issues
// instantiating the client are encountered
func NewWithClient(baseURL string, client *http.Client) (*IonClient, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("ionic: client initialization: %v", err.Error())
	}

	ic := &IonClient{
		baseURL: u,
		client:  client,
	}

	return ic, nil
}

// Delete takes an endpoint, token, params, and headers to pass as a delete call to the
// API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Delete(endpoint, token string, params *url.Values, headers http.Header) (json.RawMessage, error) {
	return requests.Delete(*ic.client, ic.baseURL, endpoint, token, params, headers)
}

// Head takes an endpoint, token, params, headers, and pagination params to pass as a
// head call to the API.  It will return any errors it encounters with the API.
func (ic *IonClient) Head(endpoint, token string, params *url.Values, headers http.Header, page *pagination.Pagination) error {
	return requests.Head(*ic.client, ic.baseURL, endpoint, token, params, headers, page)
}

// Get takes an endpoint, token, params, headers, and pagination params to pass as a
// get call to the API.  It will return a json RawMessage for the response and
// any errors it encounters with the API.
func (ic *IonClient) Get(endpoint, token string, params *url.Values, headers http.Header, page *pagination.Pagination) (json.RawMessage, error) {
	return requests.Get(*ic.client, ic.baseURL, endpoint, token, params, headers, page)
}

// Post takes an endpoint, token, params, payload, and headers to pass as a post call
// to the API.  It will return a json RawMessage for the response and any errors
// it encounters with the API.
func (ic *IonClient) Post(endpoint, token string, params *url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	return requests.Post(*ic.client, ic.baseURL, endpoint, token, params, payload, headers)
}

// Put takes an endpoint, token, params, payload, and headers to pass as a put call to
// the request.Put function.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Put(endpoint, token string, params *url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	return requests.Put(*ic.client, ic.baseURL, endpoint, token, params, payload, headers)
}
