package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuesNotes(t *testing.T) {
	ts, gitlab := Stub("stubs/notes/index.json")
	notes, err := gitlab.IssuesNotes(1, 1)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(notes), 2)
	defer ts.Close()
}

func TestIssuesNote(t *testing.T) {
	ts, gitlab := Stub("stubs/notes/single.json")
	note, err := gitlab.IssuesNote(1, 1, 1)

	assert.Equal(t, err, nil)
	assert.Equal(t, note.ID, 302)
	assert.Equal(t, note.Body, "_Status changed to closed_")
	assert.Equal(t, note.Author.ID, 1)
	assert.Equal(t, note.CreatedAt, "2013-10-02T09:22:45Z")
	defer ts.Close()
}
