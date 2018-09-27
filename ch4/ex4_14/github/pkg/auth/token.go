package auth

import (
	"context"
	"os"
)

const GithubEnvVar = "GITHUB_TOKEN"

// NewContext returns a context having the environment variable GITHUB_TOKEN.
func NewContext(ctx context.Context) context.Context {
	token := os.Getenv(GithubEnvVar)
	if token != "" {
		return context.WithValue(ctx, GithubEnvVar, token)
	}
	return ctx
}
