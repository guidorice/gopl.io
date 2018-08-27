/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/guidorice/gopl.io/ch4/ex4_12/xkcd"
)

var updateIndex = false
var query = ""

func init() {
	flag.BoolVar(&updateIndex, "update", false, "update index")
	flag.BoolVar(&updateIndex, "u", false, "update index (shorthand)")
	flag.StringVar(&query, "query", "", "query terms")
	flag.StringVar(&query, "q", "", "query terms (shorthand)")
}

func checkFlags() {
	if !updateIndex && query == "" {
		flag.Usage()
		os.Exit(2)
	}
}

func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	flag.Parse()
	checkFlags()
	indexPtr, err := xkcd.CreateOrReadIndex("")
	if err != nil {
		exitWithError(err)
	}
	if updateIndex {
		fmt.Println("updating index...")
		err := indexPtr.Update(0)
		if err != nil {
			exitWithError(err)

		}
		err = indexPtr.Save("")
		if err != nil {
			exitWithError(err)

		}
	}
	if query != "" {
		// prefix with case insensitive flag
		re, err := regexp.Compile("(?i)" + query)
		if err != nil {
			exitWithError(err)
		}
		results := indexPtr.Search(re)
		fmt.Printf("%d comics match your query (%s)\n", len(results), query)
		for _, comic := range results {
			xkcd.SearchReport.Execute(os.Stdout, comic)
		}
	} else {
		// no search query, just print the last comic in index
		comic := indexPtr.PreviousComic()
		comicJson, err := json.MarshalIndent(comic, "", "  ")
		if err != nil {
			exitWithError(err)
		}
		fmt.Printf("most recent xkcd in index:\n%s\n", string(comicJson))
	}
	fmt.Printf("%d comics in index\n", len(indexPtr.Comics))
}
