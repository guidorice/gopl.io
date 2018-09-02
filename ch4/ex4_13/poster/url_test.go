/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package poster

import (
	"net/url"
	"testing"
)

func TestDataUrl(t *testing.T) {
	want, err := url.Parse("http://www.omdbapi.com/?apikey=alex")
	if err != nil {
		t.Error(err)
	}
	got := DataUrl("alex")
	if want.String() != got.String() {
		t.Errorf("DataUrl: got %s, want %s", got, want)
	}
}

func TestPosterUrl(t *testing.T) {
	want, err := url.Parse("http://img.omdbapi.com/?apikey=alex")
	if err != nil {
		t.Error(err)
	}
	got := PosterUrl("alex")
	if want.String() != got.String() {
		t.Errorf("PosterUrl: got %s, want %s", got, want)
	}
}
