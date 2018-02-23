package events

import (
	"encoding/json"
	"fmt"
)

var validDeliveryEventActions = map[string]string{
	"delivery_failed":   "delivery_failed",
	"delivery_finished": "delivery_finished",
}

// DeliveryEventAction represents possible actions related to a delivery event
type DeliveryEventAction string

// UnmarshalJSON is a custom unmarshaller for validating a DeliveryEventAction
// or it returns an error
func (d *DeliveryEventAction) UnmarshalJSON(b []byte) error {
	var aStr string
	err := json.Unmarshal(b, &aStr)
	if err != nil {
		return err
	}

	_, ok := validDeliveryEventActions[aStr]
	if !ok {
		return fmt.Errorf("invalid delivery event action")
	}

	*d = DeliveryEventAction(validDeliveryEventActions[aStr])
	return nil
}

// MarshalJSON is a custom marshaller for validating a DeliveryEventAction
// or it returns an error
func (d DeliveryEventAction) MarshalJSON() ([]byte, error) {
	_, ok := validDeliveryEventActions[string(d)]
	if !ok {
		return nil, fmt.Errorf("invalid delivery event action")
	}

	bytes, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

//DeliveryEvent identifies the result of an delivery of a project
type DeliveryEvent struct {
	Action    DeliveryEventAction `json:"action"`
	Analysis  string              `json:"analysis"`
	ProjectID string              `json:"project_id"`
	TeamID    string              `json:"team_id"`
	Timestamp string              `json:"timestamp"`
	URL       string              `json:"url"`
}
