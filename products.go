package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/products"
)

const (
	getProductEndpoint    = "v1/vulnerability/getProducts"
	productSearchEndpoint = "v1/product/search"
)

// GetProducts takes a product ID search string and token.  It returns the product found,
// and any API errors it may encounters.
func (ic *IonClient) GetProducts(idSearch, token string) ([]products.Product, error) {
	params := &url.Values{}
	params.Set("external_id", idSearch)

	b, err := ic.Get(getProductEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get raw product: %v", err.Error())
	}

	var ps []products.Product
	err = json.Unmarshal(b, &ps)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %v", err.Error())
	}

	return ps, nil
}

// ProductSearch takes a search query. It returns a new raw json message
// of all the matching products in the Bunsen dependencies table
func (ic *IonClient) ProductSearch(searchInput products.ProductSearchQuery, token string) ([]products.Product, error) {
	if !searchInput.IsValid() {
		return nil, fmt.Errorf("Product search request not valid")
	}
	bodyBytes, err := json.Marshal(searchInput)
	if err != nil {
		// log
		return nil, err
	}
	buffer := bytes.NewBuffer(bodyBytes)
	b, err := ic.Post(productSearchEndpoint, token, nil, *buffer, nil)
	if err != nil {
		// log
		return nil, err
	}
	var ps []products.Product
	err = json.Unmarshal(b, &ps)
	if err != nil {
		// log
		return nil, err
	}
	return ps, nil
}

// GetRawProducts takes a product ID search string and token.  It returns a raw json
// message of the product found, and any API errors it may encounters.
func (ic *IonClient) GetRawProducts(idSearch, token string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("external_id", idSearch)

	b, err := ic.Get(getProductEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get raw product: %v", err.Error())
	}

	return b, nil
}

// GetProductSearch takes a search query. It returns a new raw json message of
// all the matching products in the Bunsen dependencies table
func (ic *IonClient) GetProductSearch(query, token string) ([]products.Product, error) {
	params := &url.Values{}
	params.Set("q", query)

	b, err := ic.Get(productSearchEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to GetProductSearch: %v", err.Error())
	}
	var products []products.Product
	err = json.Unmarshal(b, &products)
	if err != nil {
		return nil, fmt.Errorf("failed to parse products: %v", err.Error())
	}
	return products, nil
}
