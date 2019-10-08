package portfolios

const (
	// VulnerabilityStatsEndpoint is a string representation of the current endpoint for getting vulnerability statistics
	VulnerabilityStatsEndpoint = "v1/animal/getVulnerabilityStats"
	// VulnerabilityListEndpoint is a string representation for getting a list of vulnerabilities by type.
	VulnerabilityListEndpoint = "v1/animal/getVulnerabilityList"
	// VulnerabilityMetricsEndpoint is a string representation for getting a list of vulnerability metrics.
	VulnerabilityMetricsEndpoint = "v1/animal/getScanMetrics"
)

// VulnerabilityStat represents the vulnerabiity stat summary for the portfolio page
type VulnerabilityStat struct {
	TotalVulnerabilities      int    `json:"total_vulnerabilities"`
	UniqueVulnerabilities     int    `json:"unique_vulnerabilities"`
	MostFrequentVulnerability string `json:"most_frequent_vulnerability"`
}

// VulnerabilityListParams represents the vulnerability list paramaters
type VulnerabilityListParams struct {
	Ids      []string `json:"ids"`
	ListType string   `json:"list_type"`
	Limit    string   `json:"limit"`
}

// MetricsBody represents the metrics body
type MetricsBody struct {
	Metric     string   `json:"metric"`
	ProjectIDs []string `json:"project_ids"`
}
