package reports

import (
	"time"

	"github.com/ion-channel/ionic/analyses"
	"github.com/ion-channel/ionic/projects"
	"github.com/ion-channel/ionic/rulesets"
	"github.com/ion-channel/ionic/scanner"
)

const (
	// ReportGetProjectReportEndpoint is a string representation of the current endpoint for getting project report
	ReportGetProjectReportEndpoint = "v1/report/getProject"
	// ReportGetProjectsReportEndpoint is a string representation of the current endpoint for getting project report
	ReportGetProjectsReportEndpoint = "v1/report/getProjects"
	//ReportGetScanReportEndpoint is a string representation of the current endpoint for getting scan report
	ReportGetScanReportEndpoint = "v1/report/getScan"
	// ReportGetExportedDataEndpoint is a string representation of the current endpoint for exporting projects data
	ReportGetExportedDataEndpoint = "v1/report/getExportedData"

	// ProjectStatusErrored denotes a request for analysis has errored during
	// the run, the message field will have more details
	ProjectStatusErrored = "errored"
	// ProjectStatusPassing denotes a request for analysis has been
	// completed and is passing
	ProjectStatusPassing = "passing"
	// ProjectStatusFailing denotes a request for analysis has failed to
	// run, the message field will have more details
	ProjectStatusFailing = "failing"
	// ProjectStatusQueued denotes a request for analysis has been
	// accepted and has not begun
	ProjectStatusQueued = "queued"
	// ProjectStatusAnalyzing denotes a request for analysis has been
	// accepted and has begun
	ProjectStatusAnalyzing = "analyzing"
	// ProjectStatusPending denotes a request for analysis has not been
	// created
	ProjectStatusPending = "pending"
)

// ProjectReport gives the details of a project including past analyses
type ProjectReport struct {
	*projects.Project
	RulesetName       string             `json:"ruleset_name"`
	AnalysisSummaries []analyses.Summary `json:"analysis_summaries"`
}

// ExportedData represents the exported data
type ExportedData struct {
	CreatedAt time.Time      `json:"created_at"`
	Projects  []AnalysisData `json:"projects"`
}

// AnalysisData represents the project data
type AnalysisData struct {
	ProjectName   string `json:"project_name"`
	ProjectID     string `json:"project_id"`
	CurrentStatus string `json:"current_status"`
	VulnCount     int    `json:"vuln_count"`
	CritVulnCount int    `json:"critical_vuln_count"`
	HighVulnCount int    `json:"high_vuln_count"`
	VirusCount    int    `json:"virus_count"`
}

// NewProjectReport takes a project and analysis summaries to return a
// constructed Project Report
func NewProjectReport(project *projects.Project, summaries []analyses.Summary) *ProjectReport {
	return &ProjectReport{
		Project:           project,
		AnalysisSummaries: summaries,
	}
}

// ProjectReports is used for getting a high level overview, returning a single
// analysis
type ProjectReports struct {
	*projects.Project
	RulesetName     string            `json:"ruleset_name"`
	AnalysisSummary *analyses.Summary `json:"analysis_summary"`
	Status          string            `json:"status"`
}

// NewProjectReportsInput contains the structs needed to create a new
// ProjectReports struct
type NewProjectReportsInput struct {
	Project        *projects.Project
	Summary        *analyses.Summary
	AppliedRuleset *rulesets.AppliedRulesetSummary
	AnalysisStatus *scanner.AnalysisStatus
}

// NewProjectReports takes a project, analysis summary, and applied ruleset to
// create a summarized, high level report of a singular project. It returns this
// as a ProjectReports type.
func NewProjectReports(input NewProjectReportsInput) *ProjectReports {
	project := input.Project
	summary := input.Summary
	appliedRuleset := input.AppliedRuleset
	analysisStatus := input.AnalysisStatus

	projectStatus := ProjectStatusPending
	rulesetName := "N/A"
	if appliedRuleset != nil && appliedRuleset.RuleEvaluationSummary != nil {
		rulesetName = appliedRuleset.RuleEvaluationSummary.RulesetName
	}

	if analysisStatus != nil {
		if summary != nil {
			summary.AnalysisID = summary.ID
			summary.RulesetName = rulesetName
			summary.Trigger = "source commit"

			risk := "high"
			passed := false

			if appliedRuleset != nil {
				risk, passed = appliedRuleset.SummarizeEvaluation()
			}

			summary.Risk = risk
			summary.Passed = passed
		}

		switch analysisStatus.Status {
		case scanner.AnalysisStatusErrored:
			projectStatus = ProjectStatusErrored
		case scanner.AnalysisStatusFinished:
			projectStatus = ProjectStatusFailing
			if summary.Passed {
				projectStatus = ProjectStatusPassing
			}
		case scanner.AnalysisStatusQueued:
			projectStatus = ProjectStatusQueued
		case scanner.AnalysisStatusAnalyzing:
			projectStatus = ProjectStatusAnalyzing
		}
	}

	pr := &ProjectReports{
		Project:         project,
		RulesetName:     rulesetName,
		AnalysisSummary: summary,
		Status:          projectStatus,
	}

	return pr
}
