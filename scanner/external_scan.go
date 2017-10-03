package scanner

import (
	"encoding/json"
)

type Source struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ExternalScan struct {
	Coverage        *ExternalCoverage        `json:"coverage,omitempty"`
	Vulnerabilities *ExternalVulnerabilities `json:"vulnerabilities,omitempty"`
	Source          Source                   `json:"source"`
	Notes           string                   `json:"notes"`
	Raw             json.RawMessage          `json:"raw"`
}
