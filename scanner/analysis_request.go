package scanner

// AnalyzeRequest represents the body to send when requesting an analysis be
// done on a project.
type AnalyzeRequest struct {
	TeamID    string `json:"team_id,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
	Branch    string `json:"branch,omitempty"`
}
