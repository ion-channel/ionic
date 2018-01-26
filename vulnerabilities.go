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

	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/vulnerabilities"
)

const (
	getVulnerabilitiesEndpoint       = "v1/vulnerability/getVulnerabilities"
	getVulnerabilitiesInFileEndpoint = "v1/vulnerability/getVulnerabilitiesInFile"
	getVulnerabilityEndpoint         = "v1/vulnerability/getVulnerability"
	postVulnerabilityEndpoint        = "v1/internal/vulnerability/addVulnerability"
)

// AddVulnerability takes in a vulnerability object populated with the desired
// data to send to the API and a token to use. It will return the inserted
// vulnerability and any errors it encounters with the API.
func (ic *IonClient) AddVulnerability(newVuln *vulnerabilities.Vulnerability, token string) (*vulnerabilities.Vulnerability, error) {
	nv, err := json.Marshal(newVuln)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal new vuln into payload: %v", err.Error())
	}

	p := bytes.NewBuffer(nv)

	b, err := ic.Post(postVulnerabilityEndpoint, token, nil, *p, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to add vulnerability: %v", err.Error())
	}

	var v vulnerabilities.Vulnerability
	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json into vuln: %v", err.Error())
	}

	return &v, nil
}

// GetVulnerabilities returns a slice of Vulnerability for a given product and
// version string over a specified pagination range.  If version is left blank,
// it will not be considered in the search query.  An error is returned for
// client communication and unmarshalling errors.
func (ic *IonClient) GetVulnerabilities(product, version, token string, page *pagination.Pagination) ([]vulnerabilities.Vulnerability, error) {
	params := &url.Values{}
	params.Set("product", product)
	if version != "" {
		params.Set("version", version)
	}

	b, err := ic.Get(getVulnerabilitiesEndpoint, token, params, nil, page)
	if err != nil {
		return nil, fmt.Errorf("failed to get vulnerabilities: %v", err.Error())
	}

	var vulns []vulnerabilities.Vulnerability
	err = json.Unmarshal(b, &vulns)
	if err != nil {
		return nil, fmt.Errorf("cannot parse vulnerabilities: %v", err.Error())
	}

	return vulns, nil
}

// GetVulnerabilitiesInFile takes the location of a dependency file and returns
// a slice of vulnerabilities found for the list of dependencies.  An error is
// returned if the file can't be cannot be read, the API returns an error, or
// marshalling issues.
func (ic *IonClient) GetVulnerabilitiesInFile(filePath, token string) ([]vulnerabilities.Vulnerability, error) {
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

	b, err := ic.Post(getVulnerabilitiesInFileEndpoint, token, nil, *buff, h)
	if err != nil {
		return nil, fmt.Errorf("failed to get vulnerabilities: %v", err.Error())
	}

	var vulns []vulnerabilities.Vulnerability
	err = json.Unmarshal(b, &vulns)
	if err != nil {
		return nil, fmt.Errorf("cannot parse vulnerabilities: %v", err.Error())
	}

	return vulns, nil
}

// GetVulnerability takes an ID string and returns the vulnerability found for
// that ID.  An error is returned for API errors and marshalling errors.
func (ic *IonClient) GetVulnerability(id, token string) (*vulnerabilities.Vulnerability, error) {
	params := &url.Values{}
	params.Set("external_id", id)

	b, err := ic.Get(getVulnerabilityEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get vulnerability: %v", err.Error())
	}

	var vuln vulnerabilities.Vulnerability
	err = json.Unmarshal(b, &vuln)
	if err != nil {
		return nil, fmt.Errorf("cannot parse vulnerability: %v", err.Error())
	}

	return &vuln, nil
}

// GetRawVulnerability takes an ID string and returns the raw json message
// found for that ID.  An error is returned for API errors.
func (ic *IonClient) GetRawVulnerability(id, token string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("external_id", id)

	b, err := ic.Get(getVulnerabilityEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get vulnerability: %v", err.Error())
	}

	return b, nil
}
