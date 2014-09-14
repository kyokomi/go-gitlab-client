package gogitlab

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const (
	issuesURL        = "/issues/"                       // Get a specific issues
	projectIssuesURL = "/projects/:id/issues"           // Get a specific issues / Post a create issues
	projectIssueURL  = "/projects/:id/issues/:issue_id" // Get a specific one issue
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
	Author      User      `json:"author"`
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

	url := g.ResourceUrl(projectIssuesURL, map[string]string{":id": strconv.Itoa(projectID)})
	url += fmt.Sprintf("&page=%d", pageNo)

	var issues []*Issue

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &issues)
	}

	return issues, err
}

// ProjectIssues Get a one project issues by the authenticated user.
func (g *Gitlab) ProjectIssue(projectID, issueID int) (*Issue, error) {

	params := map[string]string{
		":id": strconv.Itoa(projectID),
		":issue_id": strconv.Itoa(issueID),
	}

	url := g.ResourceUrl(projectIssueURL, params)

	var issue *Issue

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &issue)
	}

	return issue, err
}

// ProjectCreateIssues Post a project issue by the authenticated user.
func (g *Gitlab) ProjectCreateIssues(projectID int, data []byte) ([]byte, error) {

	params := map[string]string{":id": strconv.Itoa(projectID)}

	url := g.ResourceUrl(projectIssuesURL, params)

	res, err := g.buildAndExecRequest("POST", url, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ProjectEditIssues Put a project issue by the authenticated user.
func (g *Gitlab) ProjectEditIssues(projectID, issueID int, data []byte) ([]byte, error) {

	params := map[string]string{
		":id": strconv.Itoa(projectID),
		":issue_id": strconv.Itoa(issueID),
	}

	url := g.ResourceUrl(projectIssueURL, params)

	res, err := g.buildAndExecRequest("PUT", url, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

