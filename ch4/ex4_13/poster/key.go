/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package poster

import (
	"fmt"
	"os"
)

// OmdbKey is the environment variable expected to contain the api Key from
//  http://www.omdbapi.com (required)
const omdbKeyEnvVar = "OMDB_KEY"

// Key returns the environment variable OMDB_KEY, or an error if it's missing
// or empty.
func Key() (string, error) {
	key, ok := os.LookupEnv(omdbKeyEnvVar)
	if !ok {
		return "", fmt.Errorf("environment variable %s is missing", omdbKeyEnvVar)
	}
	if key == "" {
		return "", fmt.Errorf("environment variable %s is empty", omdbKeyEnvVar)
	}
	return key, nil
}
