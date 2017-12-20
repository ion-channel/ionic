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

	"github.com/ion-channel/ionic/pagination"
)

const (
	maxIdleConns        = 25
	maxIdleConnsPerHost = 25
	maxPagingLimit      = 100
)

// IonClient represnets a communication layer with the Ion Channel API
type IonClient struct {
	baseURL     *url.URL
	bearerToken string
	client      *http.Client
}

// New takes the credentials required to talk to the API, and the base URL of
// the API and returns a client for talking to the API and an error if any
// issues instantiating the client are encountered
func New(secret, baseURL string) (*IonClient, error) {
	c := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnsPerHost,
			MaxIdleConns:        maxIdleConns,
		},
	}

	return NewWithClient(secret, baseURL, c)
}

// NewWithClient takes the credentials required to talk to the API, the base
// URL of the API, and an existing HTTP client, and returns a client for talking
// to the API and an error if any issues instantiating the client are encountered
func NewWithClient(secret, baseURL string, client *http.Client) (*IonClient, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("cannot instantiate new ion client: %v", err.Error())
	}

	ic := &IonClient{
		baseURL:     u,
		bearerToken: secret,
		client:      client,
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

// SetBearerToken takes a bearer token string to apply to the client
func (ic *IonClient) SetBearerToken(bearerToken string) {
	ic.bearerToken = bearerToken
}

func (ic *IonClient) do(method, endpoint string, params *url.Values, payload bytes.Buffer, headers http.Header, page *pagination.Pagination) (json.RawMessage, error) {
	if page == nil || page.Limit > 0 {
		ir, err := ic._do(method, endpoint, params, payload, headers, page)
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
		ir, err := ic._do(method, endpoint, params, payload, headers, page)
		if err != nil {
			return nil, fmt.Errorf("trouble paging from API: %v", err.Error())
		}
		data = append(data, ir.Data[1:len(ir.Data)-1]...)
		data = append(data, []byte(",")...)
		page.Up()
		total = ir.Meta.TotalCount
	}

	data = append(data[:len(data)-1], []byte("]")...)
	return data, nil
}

func (ic *IonClient) _do(method, endpoint string, params *url.Values, payload bytes.Buffer, headers http.Header, page *pagination.Pagination) (*IonResponse, error) {
	u := ic.createURL(endpoint, params, page)

	req, err := http.NewRequest(strings.ToUpper(method), u.String(), &payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err.Error())
	}

	if headers != nil {
		req.Header = headers
	}

	if ic.bearerToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", ic.bearerToken))
	}

	resp, err := ic.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed http request: %v", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err.Error())
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("bad response from API: %v", resp.Status)
	}

	var ir IonResponse
	err = json.Unmarshal(body, &ir)
	if err != nil {
		return nil, fmt.Errorf("malformed response: %v", err.Error())
	}

	return &ir, nil
}

// Delete takes an endpoint, params, and headers to pass as a delete call to the
// API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Delete(endpoint string, params *url.Values, headers http.Header) (json.RawMessage, error) {
	return ic.do("DELETE", endpoint, params, bytes.Buffer{}, headers, nil)
}

// Get takes an endpoint, params, headers, and pagination params to pass as a
// get call to the API.  It will return a json RawMessage for the response and
// any errors it encounters with the API.
func (ic *IonClient) Get(endpoint string, params *url.Values, headers http.Header, page *pagination.Pagination) (json.RawMessage, error) {
	return ic.do("GET", endpoint, params, bytes.Buffer{}, headers, page)
}

// Post takes an endpoint, params, payload, and headers to pass as a post call
// to the API.  It will return a json RawMessage for the response and any errors
// it encounters with the API.
func (ic *IonClient) Post(endpoint string, params *url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	return ic.do("POST", endpoint, params, payload, headers, nil)
}

// Put takes an endpoint, params, payload, and headers to pass as a put call to
// the API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Put(endpoint string, params *url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	return ic.do("PUT", endpoint, params, payload, headers, nil)
}
