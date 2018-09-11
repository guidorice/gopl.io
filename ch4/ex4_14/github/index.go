package github

import (
	"log"
	"net/http"
	"strings"
)

/*
	if err := r.ParseForm(); err != nil {
		msg := fmt.Sprintf("error parsing request: %s\n", err)
		log.Print(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
*/
// http handler to generate index page of issues, milestones, and users
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	repo := strings.Trim(path, "/")
	log.Printf("path: %s", repo)
}
