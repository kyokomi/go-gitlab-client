package gogitlab

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const (
	issuesURL        = "/issues/"                    // Get a specific issues
	projectIssuesURL = "/projects/:projectID/issues" // Get a specific issues / Post a create issues
)

// Issue struct.
type Issue struct {
	ID          int       `json:"id"`
	LocalID     int       `json:"iid"`
	ProjectID   int       `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Labels      []string  `json:"labels"`
	Milestone   Milestone `json:"milestone"`
	Author      Person    `json:"author"`
	Assignee    Person    `json:"assignee"`
	State       string    `json:"state"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	// AccessLevel int
}

// Issues Get a list of issues by the authenticated user.
func (g *Gitlab) Issues() ([]*Issue, error) {

	url := g.ResourceUrl(issuesURL, nil)

	var issues []*Issue

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &issues)
	}

	return issues, err
}

// ProjectIssues Get a list of project issues by the authenticated user.
func (g *Gitlab) ProjectIssues(projectID int, pageNo int) ([]*Issue, error) {

	url := g.ResourceUrl(projectIssuesURL, map[string]string{":projectID": strconv.Itoa(projectID)})
	url += fmt.Sprintf("&page=%d", pageNo)

	var issues []*Issue

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &issues)
	}

	return issues, err
}

// ProjectCreateIssues Post a project issue by the authenticated user.
func (g *Gitlab) ProjectCreateIssues(projectID int, data []byte) ([]byte, error) {

	url := g.ResourceUrl(projectIssuesURL, map[string]string{":projectID": strconv.Itoa(projectID)})

	res, err := g.buildAndExecRequest("POST", url, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
