package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/products"
)

const (
	getProductEndpoint = "v1/vulnerability/getProducts"
)

// GetProduct takes a product ID.  It returns the product found, and any API
// errors it may encounters.
func (ic *IonClient) GetProduct(id string) (*products.Product, error) {
	params := &url.Values{}
	params.Set("external_id", id)

	b, err := ic.get(getProductEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get raw product: %v", err.Error())
	}

	var p products.Product
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err.Error())
	}

	return &p, nil
}

// GetRawProduct takes a product ID.  It returns a raw json message of the
// product found, and any API errors it may encounters.
func (ic *IonClient) GetRawProduct(id string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("external_id", id)

	b, err := ic.get(getProductEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get raw product: %v", err.Error())
	}

	return b, nil
}
