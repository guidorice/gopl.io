package github

import (
	"os"
	"testing"
)

func TestReadIssue(t *testing.T) {
	repo := "guidorice/gopl.io"
	token := os.Getenv(GithubEnvVar)
	issueId := "4"
	if token == "" {
		t.Fatal("missing Github token")
	}
	_, err := ReadIssue(token, repo, issueId)
	if err != nil {
		t.Errorf("CreateIssue: %v", err)
	}
}
