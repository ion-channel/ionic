package events

// Event represents a singular occurance of a change within the Ion Channel
// system that can be emmitted to trigger a notification
type Event struct {
	Vulnerability *VulnerabilityEvent `json:"vulnerability,omitempty"`
	User          *UserEvent          `json:"user,omitempty"`
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
