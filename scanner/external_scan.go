package scanner

import "encoding/json"

//Source identifies the provider of external scans
type Source struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

//ExternalScan is a representation of a scan result not performed by the Ion system
type ExternalScan struct {
	Coverage        *ExternalCoverage        `json:"coverage,omitempty"`
	Vulnerabilities *ExternalVulnerabilities `json:"vulnerabilities,omitempty"`
	Source          Source                   `json:"source"`
	Notes           string                   `json:"notes"`
	Raw             *json.RawMessage         `json:"raw,omitempty"`
}
