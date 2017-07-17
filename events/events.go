package events

// Event represents a singular occurance of a change within the Ion Channel
// system that can be emmitted to trigger a notification
type Event struct {
	Analysis      *AnalysisEvent      `json:"analysis,omitempty"`
	Vulnerability *VulnerabilityEvent `json:"vulnerability,omitempty"`
	Project       *ProjectEvent       `json:"project,omitempty"`
	User          *UserEvent          `json:"user,omitempty"`
}
