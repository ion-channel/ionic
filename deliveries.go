package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/deliveries"
)

// GetDeliveryDestinations takes a team ID, and token. It returns list of deliveres and
// an error if it receives a bad response from the API or fails to unmarshal the
// JSON response from the API.
func (ic *IonClient) GetDeliveryDestinations(teamID, token string) ([]deliveries.Destination, error) {
	params := &url.Values{}
	params.Set("id", teamID)

	b, err := ic.Get(deliveries.DeliveriesGetDestinationsEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get deliveries: %v", err.Error())
	}

	var d []deliveries.Destination
	err = json.Unmarshal(b, &d)
	if err != nil {
		return nil, fmt.Errorf("failed to get deliveries: %v", err.Error())
	}

	return d, nil
}

// DeleteDeliveryDestination takes a team ID, and token. It returns errors.
func (ic *IonClient) DeleteDeliveryDestination(destinationID, token string) error {
	params := &url.Values{}
	params.Set("id", destinationID)

	_, err := ic.Delete(deliveries.DeliveriesDeleteDestinationEndpoint, token, params, nil)
	if err != nil {
		return fmt.Errorf("failed to delete delivery destination: %v", err.Error())
	}
	return err
}

// CreateDeliveryDestinations takes *CreateDestination, and token
// It returns a *CreateDestination and error
func (ic *IonClient) CreateDeliveryDestinations(dest *deliveries.CreateDestination, token string) (*deliveries.CreateDestination, error) {
	params := &url.Values{}

	b, err := json.Marshal(dest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall destination: %v", err.Error())
	}

	b, err = ic.Post(deliveries.CreateDeleteDestinationEndpoint, token, params, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination: %v", err.Error())
	}

	var a deliveries.CreateDestination
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from create destination: %v", err.Error())
	}

	dest.AccessKey = ""
	dest.SecretKey = ""
	return &a, nil
}
