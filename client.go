// Package ionic provides a direct representation of the endpoints and objects
// within the Ion Channel API
package ionic

import (
	"fmt"
	"net/http"
	"net/url"

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

//func (ic *IonClient) funcName....etc
