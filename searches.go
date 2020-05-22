package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/ion-channel/ionic/community"
	"github.com/ion-channel/ionic/products"
	"github.com/ion-channel/ionic/responses"
)

const (
	searchEndpoint = "v1/search"
)

// SearchValues structure for hold multiple search response types
type SearchValues struct {
	// Requires common fields to be explicitly
	// defined here
	Name       string    `json:"name"`
	UpdatedAt  time.Time `json:"updated_at"`
	Confidence float64   `json:"confidence"`

	*community.Repo
	*products.Product
}

// GetSearch takes a query to perform and a to be searched param
// a productidentifier search against the Ion API, assembling a slice of Ionic
// products.ProductSearchResponse objects
func (ic *IonClient) GetSearch(query, tbs, token string) ([]SearchValues, *responses.Meta, error) {
	params := &url.Values{}
	params.Set("q", query)
	params.Set("tbs", tbs)

	b, m, err := ic.Get(searchEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get productidentifiers search: %v", err.Error())
	}

	var results []SearchValues
	err = json.Unmarshal(b, &results)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal product search results: %v", err.Error())
	}
	return results, m, nil

}
