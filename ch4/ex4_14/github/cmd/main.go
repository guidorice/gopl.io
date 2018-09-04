/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/issue", issueHandler)
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe("localhost:8000", nil)
	log.Fatal(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request / %v %v", w, r)
}

func issueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("issue / %v %v", w, r)
}
