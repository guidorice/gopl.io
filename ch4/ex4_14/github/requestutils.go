package github

import (
	"errors"
	"fmt"
	"strings"
)

// ParseRequestPath parses a string of the form "org/repo" where org and repo
// are github organization and repository names. The string will be trimmed
// and if the string is badly formatted, an error returned.
func ParseRequestPath(path string) (err error, org string, repo string) {
	if len(path) == 0 {
		err := errors.New("parsing path failed, got empty string")
		return err, "", ""
	}
	trimmed := strings.Trim(path, "/")
	parts := strings.Split(trimmed, "/")
	if len(parts) != 2 {
		err := fmt.Errorf("parsing path failed, got len %d, want len 2",
			len(parts))
		return err, "", ""
	}
	return nil, parts[0], parts[1]
}
