package ionic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const maxPagingLimit = 100

// IonClient represnets a communication layer with the Ion Channel API
type IonClient struct {
	baseURL     *url.URL
	bearerToken string
	client      *http.Client
}

// IonResponse represents the response structure expected back from the Ion
// Channel API calls
type IonResponse struct {
	Data json.RawMessage `json:"data"`
	Meta Meta            `json:"meta"`
}

// Meta represents the metadata section of an IonResponse
type Meta struct {
	Version    string    `json:"version"`
	LastUpdate time.Time `json:"last_update"`
	TotalCount int       `json:"total_count"`
	Limit      int       `json:"limit"`
	Offset     int       `json:"offset"`
}

// Pagination represents the necessary elements for a paginated request
type Pagination struct {
	Offset int
	Limit  int
}

// AllItems is a convenience for requesting all items of a given entity
var AllItems = &Pagination{Offset: 0, Limit: -1}

// New takes the credentials required to talk to the API, and the base URL of
// the API and returns a client for talking to the API and an error if any
// issues instantiating the client are encountered
func New(secret string, baseURL string) (*IonClient, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("cannot instantiate new ion client: %v", err.Error())
	}

	ic := &IonClient{
		baseURL:     u,
		bearerToken: secret,
		client:      &http.Client{},
	}

	return ic, nil
}

func (ic *IonClient) createURL(endpoint string, params *url.Values, pagination *Pagination) *url.URL {
	u := *ic.baseURL
	u.Path = endpoint

	vals := &url.Values{}
	if params != nil {
		vals = params
	}

	if pagination != nil {
		vals.Set("offset", strconv.Itoa(pagination.Offset))
		vals.Set("limit", strconv.Itoa(pagination.Limit))
	}

	u.RawQuery = vals.Encode()
	return &u
}

func (ic *IonClient) do(method, endpoint string, params *url.Values, payload []byte, pagination *Pagination) (json.RawMessage, error) {
	if pagination == nil || pagination.Limit > 0 {
		ir, err := ic._do(method, endpoint, params, payload, pagination)
		if err != nil {
			return nil, err
		}

		return ir.Data, nil
	}

	page := &Pagination{Limit: maxPagingLimit, Offset: 0}
	var data json.RawMessage
	data = append(data, []byte("[")...)

	total := 1
	for page.Offset < total {
		ir, err := ic._do(method, endpoint, params, payload, page)
		if err != nil {
			return nil, fmt.Errorf("trouble paging from API: %v", err.Error())
		}
		data = append(data, ir.Data[1:len(ir.Data)-1]...)
		data = append(data, []byte(",")...)
		page.Offset += maxPagingLimit
		total = ir.Meta.TotalCount
	}

	data = append(data[:len(data)-1], []byte("]")...)
	return data, nil
}

func (ic *IonClient) _do(method, endpoint string, params *url.Values, payload []byte, pagination *Pagination) (*IonResponse, error) {
	u := ic.createURL(endpoint, params, pagination)

	req, err := http.NewRequest(strings.ToUpper(method), u.String(), nil)

	if ic.bearerToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", ic.bearerToken))
	}

	resp, err := ic.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed http request: %v", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("bad response from API: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err.Error())
	}

	var ir IonResponse
	err = json.Unmarshal(body, &ir)
	if err != nil {
		return nil, fmt.Errorf("malformed response: %v", err.Error())
	}

	return &ir, nil
}

func (ic *IonClient) delete(endpoint string, params *url.Values) (json.RawMessage, error) {
	return ic.do("DELETE", endpoint, params, nil, nil)
}

func (ic *IonClient) get(endpoint string, params *url.Values, pagination *Pagination) (json.RawMessage, error) {
	return ic.do("GET", endpoint, params, nil, pagination)
}

func (ic *IonClient) post(endpoint string, params *url.Values, payload []byte) (json.RawMessage, error) {
	return ic.do("POST", endpoint, params, payload, nil)
}

func (ic *IonClient) put(endpoint string, params *url.Values, payload []byte) (json.RawMessage, error) {
	return ic.do("PUT", endpoint, params, payload, nil)
}
