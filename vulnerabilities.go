package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	getVulnerabilitiesEndpoint       = "v1/vulnerability/getVulnerabilities"
	getVulnerabilitiesInFileEndpoint = "v1/vulnerability/getVulnerabilitiesInFile"
	getVulnerabilityEndpoint         = "v1/vulnerability/getVulnerability"
)

// Vulnerability represents a singular vulnerability record in the Ion Channel
// Platform
type Vulnerability struct {
	ID                          int             `json:"id"`
	ExternalID                  string          `json:"external_id"`
	Title                       string          `json:"title"`
	Summary                     string          `json:"summary"`
	Score                       string          `json:"score"`
	Vector                      string          `json:"vector"`
	AccessComplexity            string          `json:"access_complexity"`
	VulnerabilityAuthentication string          `json:"vulnerability_authentication"`
	ConfidentialityImpact       string          `json:"confidentiality_impact"`
	IntegrityImpact             string          `json:"integrity_impact"`
	AvailabilityImpact          string          `json:"availability_impact"`
	VulnerabilitySource         string          `json:"vulnerabilty_source"`
	AssessmentCheck             json.RawMessage `json:"assessment_check"`
	Scanner                     json.RawMessage `json:"scanner"`
	Recommendation              string          `json:"recommendation"`
	ModifiedAt                  time.Time       `json:"modified_at"`
	PublishedAt                 time.Time       `json:"published_at"`
	CreatedAt                   time.Time       `json:"created_at"`
	UpdatedAt                   time.Time       `json:"updated_at"`
	SourceID                    int             `json:"source_id"`
}

func (ic *IonClient) GetVulnerabilities(product, version string, pagination *Pagination) ([]Vulnerability, error) {
	params := &url.Values{}
	params.Set("product", product)
	if version != "" {
		params.Set("version", version)
	}

	b, err := ic.get(getVulnerabilitiesEndpoint, params, nil, pagination)
	if err != nil {
		return nil, fmt.Errorf("failed to get vulnerabilities: %v", err.Error())
	}

	var vulns []Vulnerability
	err = json.Unmarshal(b, &vulns)
	if err != nil {
		return nil, fmt.Errorf("cannot parse vulnerabilities: %v", err.Error())
	}

	return vulns, nil
}

func (ic *IonClient) GetVulnerabilitiesInFile(filePath string) ([]Vulnerability, error) {
	buff := &bytes.Buffer{}
	bw := multipart.NewWriter(buff)

	fw, err := bw.CreateFormFile("file", filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %v", err.Error())
	}

	fh, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err.Error())
	}
	defer fh.Close()

	_, err = io.Copy(fw, fh)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file to buffer: %v", err.Error())
	}

	h := http.Header{}
	h.Set("Content-Type", bw.FormDataContentType())
	bw.Close()

	b, err := ic.post(getVulnerabilitiesInFileEndpoint, nil, *buff, h)
	if err != nil {
		return nil, fmt.Errorf("failed to get vulnerabilities: %v", err.Error())
	}

	var vulns []Vulnerability
	err = json.Unmarshal(b, &vulns)
	if err != nil {
		return nil, fmt.Errorf("cannot parse vulnerabilities: %v", err.Error())
	}

	return vulns, nil
}

func (ic *IonClient) GetVulnerability(id string) (*Vulnerability, error) {
	params := &url.Values{}
	params.Set("external_id", id)

	b, err := ic.get(getVulnerabilityEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get vulnerability: %v", err.Error())
	}

	var vuln Vulnerability
	err = json.Unmarshal(b, &vuln)
	if err != nil {
		return nil, fmt.Errorf("cannot parse vulnerability: %v", err.Error())
	}

	return &vuln, nil
}
