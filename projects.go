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
	"strconv"
	"strings"

	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/projects"
)

// CreateProjectsResponse represents the response from the API when sending a
// list of projects to be created. It contains the details of each project
// created, and a list of any errors that were encountered.
type CreateProjectsResponse struct {
	Projects []projects.Project `json:"projects"`
	Errors   []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

//CreateProject takes a project object, teamId, and token to use. It returns the
// project stored or an error encountered by the API
func (ic *IonClient) CreateProject(project *projects.Project, teamID, token string) (*projects.Project, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)

	b, err := json.Marshal(project)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall project: %v", err.Error())
	}

	b, err = ic.Post(projects.CreateProjectEndpoint, token, params, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %v", err.Error())
	}

	var p projects.Project
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from create: %v", err.Error())
	}

	fields, err := p.Validate(ic.client, ic.baseURL, token)
	if err != nil {
		var errs []string
		for _, msg := range fields {
			errs = append(errs, msg)
		}
		return nil, fmt.Errorf("%v: %v", projects.ErrInvalidProject, strings.Join(errs, ", "))
	}

	return &p, nil
}

// CreateProjectsFromCSV takes a csv file location, team ID, and token to send
// the specified file to the API. All projects that are able to be created will
// be with their info returned, and a list of any errors encountered during the
// process.
func (ic *IonClient) CreateProjectsFromCSV(csvFile, teamID, token string) (*CreateProjectsResponse, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	fw, err := w.CreateFormFile("file", csvFile)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %v", err.Error())
	}

	fh, err := os.Open(csvFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err.Error())
	}

	_, err = io.Copy(fw, fh)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file contents: %v", err.Error())
	}

	w.Close()

	h := http.Header{}
	h.Set("Content-Type", w.FormDataContentType())

	b, err := ic.Post(projects.CreateProjectsFromCSVEndpoint, token, params, buf, h)
	if err != nil {
		return nil, fmt.Errorf("failed to create projects: %v", err.Error())
	}

	var resp CreateProjectsResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err.Error())
	}

	return &resp, nil
}

// GetProject takes a project ID, team ID, and token. It returns the project and
// an error if it receives a bad response from the API or fails to unmarshal the
// JSON response from the API.
func (ic *IonClient) GetProject(id, teamID, token string) (*projects.Project, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)

	b, err := ic.Get(projects.GetProjectEndpoint, token, params, nil, nil)
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

	b, err := ic.Get(projects.GetProjectEndpoint, token, params, nil, nil)
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

	b, err := ic.Get(projects.GetProjectsEndpoint, token, params, nil, page)
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

	b, err := ic.Get(projects.GetProjectByURLEndpoint, token, params, nil, nil)
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

	fields, err := project.Validate(ic.client, ic.baseURL, token)
	if err != nil {
		var errs []string
		for _, msg := range fields {
			errs = append(errs, msg)
		}
		return nil, fmt.Errorf("%v: %v", projects.ErrInvalidProject, strings.Join(errs, ", "))
	}

	params.Set("id", *project.ID)
	params.Set("team_id", *project.TeamID)

	params.Set("name", *project.Name)
	params.Set("type", *project.Type)
	params.Set("active", strconv.FormatBool(project.Active))
	params.Set("source", *project.Source)
	params.Set("branch", *project.Branch)
	params.Set("description", *project.Description)
	params.Set("ruleset_id", *project.RulesetID)
	params.Set("chat_channel", project.ChatChannel)
	params.Set("should_monitor", strconv.FormatBool(project.Monitor))

	b, err := json.Marshal(project)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall project: %v", err.Error())
	}

	b, err = ic.Put(projects.UpdateProjectEndpoint, token, params, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update projects: %v", err.Error())
	}

	var p projects.Project
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from update: %v", err.Error())
	}

	return &p, nil
}
