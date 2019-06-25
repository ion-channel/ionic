// Package ionic provides a direct representation of the endpoints and objects
// within the Ion Channel API
package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/ion-channel/ionic/errors"
	"github.com/ion-channel/ionic/pagination"
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

func (ic *IonClient) createURL(endpoint string, params *url.Values, page *pagination.Pagination) *url.URL {
	u := *ic.baseURL
	u.Path = endpoint

	vals := &url.Values{}
	if params != nil {
		vals = params
	}

	if page != nil {
		page.AddParams(vals)
	}

	u.RawQuery = vals.Encode()
	return &u
}

func (ic *IonClient) do(method, endpoint, token string, params *url.Values, payload bytes.Buffer, headers http.Header, page *pagination.Pagination) (json.RawMessage, error) {
	if page == nil || page.Limit > 0 {
		ir, err := ic._do(method, endpoint, token, params, payload, headers, page)
		if err != nil {
			return nil, err
		}

		return ir.Data, nil
	}

	page = pagination.New(0, maxPagingLimit)
	var data json.RawMessage
	data = append(data, []byte("[")...)

	total := 1
	for page.Offset < total {
		ir, err := ic._do(method, endpoint, token, params, payload, headers, page)
		if err != nil {
			err.Prepend("api: paging")
			return nil, err
		}

		data = append(data, ir.Data[1:len(ir.Data)-1]...)
		data = append(data, []byte(",")...)
		page.Up()
		total = ir.Meta.TotalCount
	}

	data = append(data[:len(data)-1], []byte("]")...)
	return data, nil
}

func (ic *IonClient) _do(method, endpoint, token string, params *url.Values, payload bytes.Buffer, headers http.Header, page *pagination.Pagination) (*IonResponse, *errors.IonError) {
	u := ic.createURL(endpoint, params, page)

	req, err := http.NewRequest(strings.ToUpper(method), u.String(), &payload)
	if err != nil {
		return nil, errors.Errors("no body", 0, "http request: failed to create: %v", err.Error())
	}

	if headers != nil {
		req.Header = headers
	}

	if token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	}

	resp, err := ic.client.Do(req)
	if err != nil {
		return nil, errors.Errors("no body", 0, "http request: failed: %v", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errors("no body", resp.StatusCode, "response body: failed to read: %v", err.Error())
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.Errors(string(body), resp.StatusCode, "api: error response")
	}

	if strings.ToUpper(method) == "HEAD" {
		return &IonResponse{}, nil
	}

	var ir IonResponse
	err = json.Unmarshal(body, &ir)
	if err != nil {
		return nil, errors.Errors(string(body), resp.StatusCode, "api: malformed response: %v", err.Error())
	}

	return &ir, nil
}

// Delete takes an endpoint, token, params, and headers to pass as a delete call to the
// API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Delete(endpoint, token string, params *url.Values, headers http.Header) (json.RawMessage, error) {
	return ic.do("DELETE", endpoint, token, params, bytes.Buffer{}, headers, nil)
}

// Get takes an endpoint, token, params, headers, and pagination params to pass as a
// get call to the API.  It will return a json RawMessage for the response and
// any errors it encounters with the API.
func (ic *IonClient) Get(endpoint, token string, params *url.Values, headers http.Header, page *pagination.Pagination) (json.RawMessage, error) {
	return ic.do("GET", endpoint, token, params, bytes.Buffer{}, headers, page)
}

// Head takes an endpoint, token, params, headers, and pagination params to pass as a
// head call to the API.  It will return any errors it encounters with the API.
func (ic *IonClient) Head(endpoint, token string, params *url.Values, headers http.Header, page *pagination.Pagination) error {
	_, err := ic.do("HEAD", endpoint, token, params, bytes.Buffer{}, headers, page)
	return err
}

// Post takes an endpoint, token, params, payload, and headers to pass as a post call
// to the API.  It will return a json RawMessage for the response and any errors
// it encounters with the API.
func (ic *IonClient) Post(endpoint, token string, params *url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	return ic.do("POST", endpoint, token, params, payload, headers, nil)
}

// Put takes an endpoint, token, params, payload, and headers to pass as a put call to
// the API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Put(endpoint, token string, params *url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	return ic.do("PUT", endpoint, token, params, payload, headers, nil)
}
