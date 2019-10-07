package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ion-channel/ionic/portfolios"
)

// GetVulnerabilityStats takes slice of project ids and token and returns vulnerability stats and any errors
func (ic *IonClient) GetVulnerabilityStats(ids []string, token string) (*portfolios.VulnerabilityStat, error) {
	p := struct {
		Ids []string `json:"ids"`
	}{
		ids,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(portfolios.VulnerabilityStatsEndpoint, token, nil, *bytes.NewBuffer(b), nil)

	if err != nil {
		return nil, fmt.Errorf("failed to request vulnerability list: %v", err.Error())
	}

	var vs portfolios.VulnerabilityStat
	err = json.Unmarshal(r, &vs)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal vunlerability stats response: %v", err.Error())
	}

	return &vs, nil
}
