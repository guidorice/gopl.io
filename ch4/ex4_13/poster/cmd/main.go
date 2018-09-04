/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"image/jpeg"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/guidorice/gopl.io/ch4/ex4_13/poster"
)

func haltIf(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	movie := poster.Movie{}
	apiKey, err := poster.Key()
	haltIf(err)
	if len(os.Args[1:]) < 1 {
		haltIf(errors.New("error: please provide a movie omdbId, title, or search string"))
	}
	query := strings.Join(os.Args[1:], " ")
	if detectId(query) {
		params := poster.Params{poster.Id: query}
		movie, err = poster.MovieRequest(apiKey, params)
		haltIf(err)
		if movie.Response != "True" {
			// try again with title query
			params := poster.Params{poster.Title: query}
			movie, err = poster.MovieRequest(apiKey, params)
			haltIf(err)
		}
	}
	if movie.ImdbID != "" {
		fetchPoster(apiKey, movie)
		os.Exit(0)
	}
	// try again with search term
	params := poster.Params{poster.Search: query}
	result, err := poster.SearchRequest(apiKey, params)
	if len(result.Search) == 0 {
		fmt.Printf("No results match your query '%s'.\n", query)
		os.Exit(0)
	}
	fmt.Print("Here are some results. please select a movie:\n\n")
	moviePtr := selectMovie(result)
	fmt.Printf("Fetching %s...\n", moviePtr.Title)
	fetchPoster(apiKey, *moviePtr)
}

// selectMovie prompts user to select from 0..n in the SearchResult, returning
// a poster.Movie.
func selectMovie(result poster.SearchResult) *poster.Movie {
	for i, movie := range result.Search {
		fmt.Printf("[%d] %s (%d)\n", i, movie.Title, movie.Year)
	}
	fmt.Printf("\nChoose a number:")
	reader := bufio.NewReader(os.Stdin)
	// shortcut relying on the fact that omdb returns 10 results max (so ReadRune
	// will work, the index is always in the range 0-9.
	char, _, err := reader.ReadRune()
	haltIf(err)
	selectedIndex, err := strconv.Atoi(string(char))
	if err != nil || selectedIndex > len(result.Search) {
		fmt.Printf(
			"Sorry, I didn't get that. Enter a number from 0 to %d\n\n",
			len(result.Search)-1,
		)
		return selectMovie(result)
	}
	return result.Search[selectedIndex]
}

// fetchPoster fetches the jpeg poster with the apiKey and movie info. It
// writes to a file and calls "open" to preview it (open is an OS X utility).
func fetchPoster(apiKey string, movie poster.Movie) {
	id := movie.ImdbID
	img, err := poster.PosterRequest(apiKey, id)
	haltIf(err)
	filename := fmt.Sprintf("poster-%s.jpg", id)
	f, err := os.Create(filename)
	haltIf(err)
	defer f.Close()
	jpeg.Encode(f, img, nil)
	fmt.Printf("wrote poster to %s\n", filename)
	cmd := exec.Command("open", filename)
	err = cmd.Run()
	haltIf(err)
}

func detectId(query string) bool {
	re := regexp.MustCompile(`\d`)
	return re.MatchString(query)
}
