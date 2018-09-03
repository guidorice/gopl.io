package github

import (
	"os"
	"testing"
)

func TestCloseIssue(t *testing.T) {
	repo := "guidorice/gopl.io"
	token := os.Getenv(GithubEnvVar)
	issueId := "16"
	if token == "" {
		t.Fatal("missing Github token")
	}
	_, err := CloseIssue(token, repo, issueId)
	if err != nil {
		t.Errorf("CloseIssue: %v", err)
	}
}
