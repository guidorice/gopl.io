package github

import (
	"os"
	"testing"
)

func TestUpdateIssue(t *testing.T) {
	var err error
	repo := "guidorice/gopl.io"
	token := os.Getenv(GithubEnvVar)
	if token == "" {
		t.Fatal("missing Github token")
	}
	issueSrc := IssueCreate{
		Title: "test issue",
		Body:  "lorem ipsum",
	}
	issue, err := CreateIssue(token, repo, issueSrc)
	if err != nil {
		t.Error(err)
	}
	want := "hello world"
	issuePatch := Issue{Number: issue.Number, Body: want}
	issue, err = UpdateIssue(token, repo, issuePatch)
	if err != nil {
		t.Error(err)
	}
	if issue.Body != want {
		t.Errorf("UpdateIssue: got %s, want %s", issue.Body, want)
	}
}
