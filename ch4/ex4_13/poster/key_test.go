/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package poster

import (
	"os"
	"testing"
)

func TestKey(t *testing.T) {
	const want = "hello999"
	os.Setenv(omdbKeyEnvVar, want)
	key, err := Key()
	if err != nil {
		t.Error(err)
	}
	if key != want {
		t.Errorf("Key() got %s, want %s", key, want)
	}
}
