package scanner

import (
	"encoding/json"
)

type Source struct {
  Name string `json:"name"`
  Url string `json:"url"`
}

type ExternalScan struct {
  Coverage  ExternalCoverage `json:"coverage"`
  Vulnerabilities  ExternalVulnerabilities `json:"vulnerabilities"`
  Source Source `json:"source"`
  Notes string `json:"notes"`
  Raw []json.RawMessage `json:"raw"`
}
