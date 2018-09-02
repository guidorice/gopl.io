/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package poster

import (
	"image"
	"testing"
)

type testCase struct {
	M Movie
	P Params
}

// testTable is a slice of test cases where P are the api query parameters, and
// M is the Movie struct expected in response.
var movieRequestCases = []testCase{
	{
		P: Params{
			Title: "Guido",
		},
		M: Movie{
			ImdbID: "tt1743334",
		},
	},
	{
		P: Params{
			Title: "Ghostbusters",
		},
		M: Movie{
			ImdbID: "tt0087332",
		},
	},
	{
		P: Params{
			Id: "tt1285016",
		},
		M: Movie{
			ImdbID: "tt1285016",
			Title:  "The Social Network",
		},
	},
}

func TestMovieRequestCases(t *testing.T) {
	key, err := Key()
	if err != nil {
		t.Error(err)
	}
	for _, testCase := range movieRequestCases {
		params := testCase.P
		want := testCase.M
		got, err := MovieRequest(key, params)
		if err != nil {
			t.Error(err)
		}
		if got.ImdbID != want.ImdbID {
			t.Errorf("MovieRequest: want %v, got %v", want, got)
		}
	}
}

type searchCase struct {
	P Params
	R SearchResult
}

var searchRequestCases = []searchCase{
	{
		P: Params{
			Search: "jaws",
		},
		R: SearchResult{
			TotalResults: 84,
			Response:     "True",
		},
	},
}

func TestSearchRequestCases(t *testing.T) {
	key, err := Key()
	if err != nil {
		t.Error(err)
	}
	for _, testCase := range searchRequestCases {
		params := testCase.P
		want := testCase.R
		got, err := SearchRequest(key, params)
		if err != nil {
			t.Error(err)
		}
		if len(got.Search) != 10 {
			t.Errorf("SearchRequest: want %v, got %v", 10, len(got.Search))
		}
		if got.Response != want.Response || got.TotalResults != want.TotalResults {
			t.Errorf("SearchRequest: want %v, got %v", want.Response, got.Response)
		}
	}
}

func TestPosterRequest(t *testing.T) {
	key, err := Key()
	if err != nil {
		t.Error(err)
	}
	id := "tt1285016"
	want := image.Rect(0, 0, 190, 300)
	img, err := PosterRequest(key, id)
	if err != nil {
		t.Error(err)
	}
	if want != img.Bounds() {
		t.Errorf("PosterRequst: want image size: %v, got %v", want, img.Bounds())
	}
}
