package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/projects"
)

const (
	createProjectEndpoint   = "v1/project/createProject"
	getProjectEndpoint      = "v1/project/getProject"
	getProjectByURLEndpoint = "v1/project/getProjectByUrl"
	getProjectsEndpoint     = "v1/project/getProjects"
	updateProjectEndpoint   = "v1/project/updateProject"
)

//CreateProject takes a project object and token to use. It returns the
// project stored or an error encountered by the API
func (ic *IonClient) CreateProject(project *projects.Project, teamID, token string) (*projects.Project, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)

	b, err := json.Marshal(project)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall project: %v", err.Error())
	}

	b, err = ic.Post(createProjectEndpoint, token, params, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %v", err.Error())
	}

	var p projects.Project
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from create: %v", err.Error())
	}

	return &p, nil
}

// GetProject takes a project ID, team ID, and token. It returns the project and
// an error if it receives a bad response from the API or fails to unmarshal the
// JSON response from the API.
func (ic *IonClient) GetProject(id, teamID, token string) (*projects.Project, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)

	b, err := ic.Get(getProjectEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %v", err.Error())
	}

	var p projects.Project
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %v", err.Error())
	}

	return &p, nil
}

// GetRawProject takes a project ID, team ID, and token. It returns the raw json of the
// project.  It also returns any API errors it may encounter.
func (ic *IonClient) GetRawProject(id, teamID, token string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)

	b, err := ic.Get(getProjectEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %v", err.Error())
	}

	return b, nil
}

// GetProjects takes a team ID and returns the projects for that team.  It
// returns an error for any API errors it may encounter.
func (ic *IonClient) GetProjects(teamID, token string, page *pagination.Pagination) ([]projects.Project, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)

	b, err := ic.Get(getProjectsEndpoint, token, params, nil, page)
	if err != nil {
		return nil, fmt.Errorf("failed to get projects: %v", err.Error())
	}

	var pList []projects.Project
	err = json.Unmarshal(b, &pList)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal projects: %v", err.Error())
	}

	return pList, nil
}

// GetProjectByURL takes a uri, teamID, and API token to request the noted
// project from the API. It returns the project and any errors it encounters
// with the API.
func (ic *IonClient) GetProjectByURL(uri, teamID, token string) (*projects.Project, error) {
	params := &url.Values{}
	params.Set("url", uri)
	params.Set("team_id", teamID)

	b, err := ic.Get(getProjectByURLEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get projects by url: %v", err.Error())
	}

	var p projects.Project
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal projects: %v", err.Error())
	}

	return &p, nil
}

//UpdateProject takes a project to update and token to use. It returns the
// project stored or an error encountered by the API
func (ic *IonClient) UpdateProject(project *projects.Project, token string) (*projects.Project, error) {
	params := &url.Values{}

	params.Set("id", project.ID)
	params.Set("team_id", project.TeamID)

	params.Set("name", project.Name)
	params.Set("type", project.Type)
	params.Set("active", strconv.FormatBool(project.Active))
	params.Set("source", project.Source)
	params.Set("branch", project.Branch)
	params.Set("description", project.Description)
	params.Set("ruleset_id", project.RulesetID)
	params.Set("chat_channel", project.ChatChannel)
	params.Set("should_monitor", strconv.FormatBool(project.Monitor))

	b, err := ic.Put(updateProjectEndpoint, token, params, bytes.Buffer{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get projects: %v", err.Error())
	}

	var p projects.Project
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from update: %v", err.Error())
	}

	return &p, nil
}
