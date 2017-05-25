package ionic

// Event represents a singular occurance of a change within the Ion Channel
// system that can be emmitted to trigger a notification
type Event struct {
	Vulnerability VulnerabilityEvent `json:"vulnerability,omitempty"`
}

// VulnerabilityEvent represents the vulnerability releated segement of an Event within Ion Channel
type VulnerabilityEvent struct {
	Updates  []string       `json:"updates,omitempty"`
	Projects []ProjectEvent `json:"projects,omitempty"`
}

// ProjectEvent represents a project within an Event within Ion Channel
type ProjectEvent struct {
	Project  string   `json:"project"`
	Org      string   `json:"org"`
	Versions []string `json:"versions"`
}

// Append takes an event to join and leaves the union of the two events
func (e *Event) Append(toAppend Event) {
	for _, v := range toAppend.Vulnerability.Updates {
		if !e.contains(v) {
			e.Vulnerability.Updates = append(e.Vulnerability.Updates, v)
		}
	}
}

func (e *Event) contains(vuln string) bool {
	if e.Vulnerability.Updates == nil {
		return false
	}

	for _, existing := range e.Vulnerability.Updates {
		if existing == vuln {
			return true
		}
	}

	return false
}
