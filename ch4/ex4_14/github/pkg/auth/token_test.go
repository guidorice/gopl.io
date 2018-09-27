package auth

import (
	"context"
	"os"
	"testing"
)

func TestNewContext_WithToken(t *testing.T) {
	token := "foo999"
	os.Setenv(GithubEnvVar, token)
	ctx := context.Background()
	got := NewContext(ctx).Value(GithubEnvVar)
	if got != token {
		t.Errorf("NewContext: got %v, want %v", got, token)
	}
}

func TestNewContext_WithoutToken(t *testing.T) {
	ctx := context.Background()
	os.Unsetenv(GithubEnvVar)
	got := NewContext(ctx).Value(GithubEnvVar)
	if got != nil {
		t.Errorf("NewContext: got %v, want nil", got)
	}
}
