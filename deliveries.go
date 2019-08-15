package ionic

import (
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
