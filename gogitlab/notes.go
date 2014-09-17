package gogitlab

import (
	"strconv"
	"encoding/json"
)

const (
	issuesNotesURL = "/projects/:id/issues/:issue_id/notes"
	issuesNoteURL = "/projects/:id/issues/:issue_id/notes/:note_id"
	createIssueNoteURL = "/projects/:id/issues/:issue_id/notes"
)

// Note are comments on snippets, issues or merge requests. struct
type Note struct {
	ID        int          `json:"id"`
//	Attachment interface{} `json:"attachment"`
	Author    User         `json:"author"`
	Body      string       `json:"body"`
	CreatedAt string       `json:"created_at"`
}

// IssuesNotes GET a list of all notes for a single issue. by the authenticated user.
func (g *Gitlab) IssuesNotes(projectID, issueID int) ([]*Note, error) {

	param := map[string]string{
		":id":       strconv.Itoa(projectID),
		":issue_id": strconv.Itoa(issueID),
	}
	url := g.ResourceUrl(issuesNotesURL, param)

	var notes []*Note

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &notes)
	}

	return notes, err
}

// IssuesNote GET a single note for a specific project issue. by the authenticated user.
func (g *Gitlab) IssuesNote(projectID, issueID, noteID int) (*Note, error) {

	param := map[string]string{
		":id":       strconv.Itoa(projectID),
		":issue_id": strconv.Itoa(issueID),
		":note_id":  strconv.Itoa(noteID),
	}
	url := g.ResourceUrl(issuesNoteURL, param)

	var note *Note

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &note)
	}

	return note, err
}

// CreateIssuesNote POST a new note to a single project issue. by the authenticated user.
func (g *Gitlab) CreateIssuesNote(projectID, issueID int, data []byte) ([]byte, error) {

	param := map[string]string{
		":id":       strconv.Itoa(projectID),
		":issue_id": strconv.Itoa(issueID),
	}
	url := g.ResourceUrl(createIssueNoteURL, param)

	return g.buildAndExecRequest("POST", url, data)
}
