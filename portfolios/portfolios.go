package portfolios

const (
	// VulnerabilityStatsEndpoint is a string representation of the current endpoint for getting vulnerability statistics
	VulnerabilityStatsEndpoint = "v1/animal/getVulnerabilityStats"
	// VulnerabilityListEndpoint is a string representation for getting a list of vulnerabilities by type.
	VulnerabilityListEndpoint = "v1/animal/getVulnerabilityList"
	// VulnerabilityMetricsEndpoint is a string representation for getting a list of vulnerability metrics.
	VulnerabilityMetricsEndpoint = "v1/animal/getScanMetrics"
	// PortfolioPassFailSummaryEndpoint is a string representation for getting a portfolio status summary.
	PortfolioPassFailSummaryEndpoint = "v1/ruleset/getStatuses"
	// PortfolioStartedErroredSummaryEndpoint is a string representation for getting the started and errored count for a list of projects
	PortfolioStartedErroredSummaryEndpoint = "v1/scanner/getStatuses"
	// PortfolioGetAffectedProjectIdsEndpoint is a string representation for getting a list of affected projects.
	PortfolioGetAffectedProjectIdsEndpoint = "v1/animal/getAffectedProjectIds"
	// PortfolioGetAffectedProjectsInfoEndpoint is a string representation for getting the name and version of projects from a list of project ids
	PortfolioGetAffectedProjectsInfoEndpoint = "v1/project/getAffectedProjectsInfo"
	// DependencyStatsEndpoint is a string representation of the current endpoint for getting dependencies statistics
	DependencyStatsEndpoint = "v1/animal/getDependencyStats"
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
	ListType string   `json:"list_type,omitempty"`
	Limit    string   `json:"limit,omitempty"`
}

// MetricsBody represents the metrics body
type MetricsBody struct {
	Metric     string   `json:"metric"`
	ProjectIDs []string `json:"project_ids"`
}

// PortfolioStatusSummary represents a summary of status for projects in a
// Portfolio
type PortfolioStatusSummary struct {
	PassingProjects int `json:"passing_projects"`
	FailingProjects int `json:"failing_projects"`
	ErroredProjects int `json:"errored_projects"`
	PendingProjects int `json:"pending_projects"`
}

// PortfolioPassingFailingSummary represents a summary of passing and failing for projects
type PortfolioPassingFailingSummary struct {
	PassingProjects int `json:"passing_projects"`
	FailingProjects int `json:"failing_projects"`
}

// PortfolioStartedErroredSummary represents a summary of started and errored statuses for projects
type PortfolioStartedErroredSummary struct {
	AnalyzingProjects int `json:"analyzing_projects"`
	ErroredProjects   int `json:"errored_projects"`
	FinishedProjects  int `json:"finished_projects"`
}

// PortfolioRequestedIds represents a list of IDs to send to a request
type PortfolioRequestedIds struct {
	IDs []string `json:"ids"`
}

// AffectedProject represents a single project affected by a vulnerability
type AffectedProject struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Version         string `json:"version"`
	Vulnerabilities int    `json:"vulnerabilities"`
}

// DependencyStat represents the dependency stat summary for the portfolio page
type DependencyStat struct {
	DirectDependencies     int `json:"direct_dependencies"`
	TransitiveDependencies int `json:"transitive_dependencies"`
	OutdatedDependencies   int `json:"outdated_dependencies"`
	NoVersionSpecified     int `json:"no_vesion_dependencies"`
}
