package events

import (
	"encoding/json"
	"fmt"
)

var validDeliveryEventActions = map[string]string{
	"delivery_failed":          "delivery_failed",
	"artifact_delivery_failed": "artifact_delivery_failed",
	"report_delivery_failed":   "report_delivery_failed",
	"seva_delivery_failed":     "seva_delivery_failed",

	"delivery_finished":          "delivery_finished",
	"artifact_delivery_finished": "artifact_delivery_finished",
	"report_delivery_finished":   "report_delivery_finished",
	"seva_delivery_finished":     "seva_delivery_finished",

	"delivery_canceled":          "delivery_canceled",
	"artifact_delivery_canceled": "artifact_delivery_canceled",
	"report_delivery_canceled":   "report_delivery_canceled",
	"seva_delivery_canceled":     "seva_delivery_canceled",
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

	bytes := []byte(fmt.Sprintf("\"%v\"", d))

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
	Filename  string              `json:"filename"`
}
