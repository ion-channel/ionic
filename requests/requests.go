package requests

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
	"github.com/ion-channel/ionic/responses"
)

const (
	maxPagingLimit = 100
)

func do(client http.Client, method string, baseURL *url.URL, endpoint, token string, params *url.Values, payload bytes.Buffer, headers http.Header, page *pagination.Pagination) (json.RawMessage, error) {
	if page == nil || page.Limit > 0 {
		ir, err := _do(client, method, baseURL, endpoint, token, params, payload, headers, page)
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
		ir, err := _do(client, method, baseURL, endpoint, token, params, payload, headers, page)
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

func _do(client http.Client, method string, baseURL *url.URL, endpoint, token string, params *url.Values, payload bytes.Buffer, headers http.Header, page *pagination.Pagination) (*responses.IonResponse, *errors.IonError) {
	u := createURL(baseURL, endpoint, params, page)

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

	resp, err := client.Do(req)
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
		return &responses.IonResponse{}, nil
	}

	var ir responses.IonResponse
	err = json.Unmarshal(body, &ir)
	if err != nil {
		return nil, errors.Errors(string(body), resp.StatusCode, "api: malformed response: %v", err.Error())
	}

	return &ir, nil
}

func createURL(baseURL *url.URL, endpoint string, params *url.Values, page *pagination.Pagination) *url.URL {
	u := *baseURL
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

// Delete takes an endpoint, token, params, and headers to pass as a delete call to the
// API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func Delete(client http.Client, baseURL *url.URL, endpoint, token string, params *url.Values, headers http.Header) (json.RawMessage, error) {
	return do(client, "DELETE", baseURL, endpoint, token, params, bytes.Buffer{}, headers, nil)
}

// Head takes an endpoint, token, params, headers, and pagination params to pass as a
// head call to the API.  It will return any errors it encounters with the API.
func Head(client http.Client, baseURL *url.URL, endpoint, token string, params *url.Values, headers http.Header, page *pagination.Pagination) error {
	_, err := do(client, "HEAD", baseURL, endpoint, token, params, bytes.Buffer{}, headers, page)
	return err
}

// Get takes an endpoint, token, params, headers, and pagination params to pass as a
// get call to the API.  It will return a json RawMessage for the response and
// any errors it encounters with the API.
func Get(client http.Client, baseURL *url.URL, endpoint, token string, params *url.Values, headers http.Header, page *pagination.Pagination) (json.RawMessage, error) {
	return do(client, "GET", baseURL, endpoint, token, params, bytes.Buffer{}, headers, page)
}

// Post takes an endpoint, token, params, payload, and headers to pass as a post call
// to the API.  It will return a json RawMessage for the response and any errors
// it encounters with the API.
func Post(client http.Client, baseURL *url.URL, endpoint, token string, params *url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	return do(client, "POST", baseURL, endpoint, token, params, payload, headers, nil)
}

// Put takes an endpoint, token, params, payload, and headers to pass as a put call to
// the API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func Put(client http.Client, baseURL *url.URL, endpoint, token string, params *url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	return do(client, "PUT", baseURL, endpoint, token, params, payload, headers, nil)
}
