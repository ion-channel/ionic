package scanner

import "encoding/json"

//Source needs a comment
type Source struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

//ExternalScan needs a comment
type ExternalScan struct {
	Coverage        *ExternalCoverage        `json:"coverage,omitempty"`
	Vulnerabilities *ExternalVulnerabilities `json:"vulnerabilities,omitempty"`
	Source          Source                   `json:"source"`
	Notes           string                   `json:"notes"`
	Raw             *json.RawMessage         `json:"raw,omitempty"`
}
