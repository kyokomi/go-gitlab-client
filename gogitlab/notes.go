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

type Note struct {
	ID        int          `json:"id"`
//	Attachment interface{} `json:"attachment"`
	Author    User         `json:"author"`
	Body      string       `json:"body"`
	CreatedAt string       `json:"created_at"`
}

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
