/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/guidorice/gopl.io/ch4/ex4_14/github"
)

func main() {
	http.HandleFunc("/issue", issueHandler)
	http.HandleFunc("/", github.IndexHandler)
	err := http.ListenAndServe("localhost:8000", nil)
	log.Fatal(err)
}

func issueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("issue / %v %v", w, r)
}
