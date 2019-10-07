package portfolios

const (
	// VulnerabilityStatsEndpoint is a string representation of the current endpoint for getting vulnerability statistics
	VulnerabilityStatsEndpoint = "v1/animal/getVulnerabilityStats"
)

// VulnerabilityStat represents the vulnerabiity stat summary for the portfolio page
type VulnerabilityStat struct {
	TotalVulnerabilities      int    `json:"total_vulnerabilities"`
	UniqueVulnerabilities     int    `json:"unique_vulnerabilities"`
	MostFrequentVulnerability string `json:"most_frequent_vulnerability"`
}
