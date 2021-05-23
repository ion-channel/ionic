package events

// Event represents a singular occurrence of a change within the Ion Channel
// system that can be emitted to trigger a notification
type Event struct {
	Analysis      *AnalysisEvent      `json:"analysis,omitempty"`
	Delivery      *DeliveryEvent      `json:"delivery,omitempty"`
	Project       *ProjectEvent       `json:"project,omitempty"`
	Vulnerability *VulnerabilityEvent `json:"vulnerability,omitempty"`
	User          *UserEvent          `json:"user,omitempty"`
}
