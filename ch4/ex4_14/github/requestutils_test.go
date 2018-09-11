package github

import (
	"errors"
	"reflect"
	"testing"
)

var requestPathTests = []struct {
	arg  string
	err  error
	org  string
	repo string
}{
	{"org/repo", nil, "org", "repo"},
	{"/org/repo", nil, "org", "repo"},
	{"/org/repo/", nil, "org", "repo"},
	{"", errors.New("parsing path failed, got empty string"), "", ""},
	{"someorg", errors.New("parsing path failed, got len 1, want len 2"), "", ""},
	{"someorg/", errors.New("parsing path failed, got len 1, want len 2"), "", ""},
	{"somedev/repoman", nil, "somedev", "repoman"},
}

func TestParseRequestPath(t *testing.T) {
	for _, test := range requestPathTests {
		err, org, repo := ParseRequestPath(test.arg)
		if !reflect.DeepEqual(err, test.err) {
			t.Errorf("ParseRequestPath: want %v, got: %v",
				test.err, err)
		}
		if org != test.org {
			t.Errorf("ParseRequestPath: want org %v, got: org %v",
				test.org, org)
		}
		if repo != test.repo {
			t.Errorf("ParseRequestPath: want repo %v, got: repo %v",
				test.repo, repo)
		}
	}
}
