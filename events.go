package ionic

// Event represents a singular occurance of a change within the Ion Channel
// system that can be emmitted to trigger a notification
type Event struct {
	Vulns EventVulnerability `json:"vulns"`
}

type EventVulnerability struct {
	Updates  []string       `json:"updates"`
	Projects []EventProject `json:"projects"`
}

type EventProject struct {
	Project  string   `json:"project"`
	Org      string   `json:"org"`
	Versions []string `json:"versions"`
}

// Append takes an event to join and leaves the union of the two events
func (e *Event) Append(toAppend Event) {
	for _, v := range toAppend.Vulns.Updates {
		if !e.contains(v) {
			e.Vulns.Updates = append(e.Vulns.Updates, v)
		}
	}
}

func (e *Event) contains(vuln string) bool {
	if e.Vulns.Updates == nil {
		return false
	}

	for _, existing := range e.Vulns.Updates {
		if existing == vuln {
			return true
		}
	}

	return false
}
