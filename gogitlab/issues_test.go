package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssues(t *testing.T) {
	ts, gitlab := Stub("stubs/issues/index.json")
	issues, err := gitlab.Issues()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(issues), 2)
	defer ts.Close()
}

func TestProjectIssues(t *testing.T) {
	ts, gitlab := Stub("stubs/issues/index.json")
	issues, err := gitlab.ProjectIssues(1, 1)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(issues), 2)
	defer ts.Close()
}

func TestProjectCreateIssues(t *testing.T) {
	ts, gitlab := Stub("")
	_, err := gitlab.ProjectCreateIssues(1, nil)

	assert.Equal(t, err, nil)
	defer ts.Close()
}
