/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package poster

// Movie is a struct for serializing omdb json records.
type Movie struct {
	Actors     string
	Awards     string
	BoxOffice  string
	Country    string
	Director   string
	Dvd        string
	Genre      string
	ImdbID     string
	ImdbRating string
	ImdbVotes  string
	Language   string
	Metascore  string
	Plot       string
	Poster     string
	Production string
	Rated      string
	Ratings    []*Rating
	Released   string
	Response   string
	Runtime    string
	Title      string
	Type       string
	Website    string
	Writer     string
	Year       int `json:",string"`
}

// Rating is a struct for serializing omdb json records.
type Rating struct {
	Source string
	Value  string
}

// Rating is struct for serializing omdb json reconds
type SearchResult struct {
	Search       []*Movie
	TotalResults int `json:",string"`
	Response     string
}
