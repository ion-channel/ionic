package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/reports"
)

const (
	reportGetReportEndpoint = "v1/report/getAnalysis"
)

func (ic *IonClient) GetReport(id, teamID, projectID string) (*reports.Report, error) {
	params := &url.Values{}
	params.Set("analysis_id", id)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.get(reportGetReportEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get report: %v", err.Error())
	}

	var r reports.Report
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, fmt.Errorf("failed to get report: %v", err.Error())
	}

	return &r, nil
}
