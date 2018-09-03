/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// ex 4.11: Build a tool that lets users create, read, update, and close github
// issues from the command line, invoking their preferred text editor when sub-
// stantial text input is required.
package main

// TODO: lower timeouts in all http clients https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
// TODO: implement updating w/ editor

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/guidorice/gopl.io/ch4/ex4_11/github"
)

type action int

const (
	undefined action = iota
	search
	create
	read
	update
	close
)

// tmp struct for flag parsing
type flags struct {
	close   int
	create  bool
	message string
	read    int
	repo    string
	search  bool
	title   string
	update  int
}

// runtime options struct (after flag parse)
type options struct {
	action  action
	issue   int
	message string
	query   string
	repo    string
	search  string
	title   string
	token   string
}

var f flags
var opts options

// package init
func init() {
	flag.BoolVar(&f.create, "create", false, "create a github issue")

	flag.IntVar(&f.close, "close", -1, "close a github issue")
	flag.IntVar(&f.read, "read", -1, "read a github issue")
	flag.IntVar(&f.update, "update", -1, "update a github issue")

	flag.BoolVar(&f.search, "search", false, "search issues by terms")
	flag.StringVar(&f.repo, "repo", "", "repository name")
	flag.StringVar(&f.title, "title", "", "title of new issue")
	flag.StringVar(&f.message, "message", "", "message body of new issue")
	flag.Usage = Usage
}

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	// TODO: add detailed usage w/ examples
	flag.PrintDefaults()
}

// checkFlags converts boolean, int, etc flags into more readable types.
func checkFlags() {
	opts.repo = f.repo
	if f.create {
		opts.action = create
		opts.title = f.title
		opts.message = f.message
		if opts.repo == "" || opts.title == "" {
			log.Printf("Example: -create -title \"title...\" -message \"message...\"")
			flag.Usage()
			os.Exit(2)
		}
	} else if f.read > 0 {
		opts.action = read
		opts.issue = f.read
	} else if f.close > 0 {
		opts.action = close
		opts.issue = f.close
		if opts.issue == 0 {
			flag.Usage()
			os.Exit(2)
		}
	} else if f.search {
		opts.action = search
		opts.query = strings.Join(flag.Args(), " ")
	} else if f.update > 0 {
		opts.action = update
		opts.issue = f.update
		opts.message = f.message
		opts.title = f.title
	}
}

func checkAuth() {
	opts.token = os.Getenv(github.GithubEnvVar)
	if opts.token == "" {
		log.Printf("warning: no auth token, set environment variable %s\n",
			github.GithubEnvVar,
		)
	}
}

const IssuesSearchTemplate = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var IssuesSearchReport = template.Must(template.New("searchResult").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(IssuesSearchTemplate))

const IssueTemplate = `issue#:  {{.Number}} {{.HTMLURL}}
title:   {{.Title}}
state:   {{.State}}
age:    {{.CreatedAt | daysAgo}} days
creator: {{.User.Login}} {{.User.HTMLURL}}

{{.Body}}
`

var IssueReport = template.Must(template.New("issue").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(IssueTemplate))

func main() {
	flag.Parse()
	checkFlags()
	checkAuth()
	switch opts.action {
	case create:
		if opts.message == "" {
			b, err := editorMessage()
			if err != nil {
				log.Fatal(err)
			}
			opts.message = string(b)
		}
		tmpl := github.IssueCreate{
			Title: opts.title,
			Body:  opts.message,
		}
		issue, err := github.CreateIssue(opts.token, opts.repo, tmpl)
		if err != nil {
			log.Fatal(err)
		}
		IssueReport.Execute(os.Stdout, issue)
	case read:
		issue, err := github.ReadIssue(opts.token, opts.repo, string(opts.issue))
		if err != nil {
			log.Fatal(err)
		}
		IssueReport.Execute(os.Stdout, issue)
	case close:
		issue, err := github.CloseIssue(opts.token, opts.repo, string(opts.issue))
		if err != nil {
			log.Fatal(err)
		}
		IssueReport.Execute(os.Stdout, issue)
	case search:
		// prefix the search terms with a single repo name
		terms := "repo:" + string(opts.repo) + " " + opts.query
		results, err := github.SearchIssues(opts.token, terms)
		if err != nil {
			log.Fatal(err)
		}
		IssuesSearchReport.Execute(os.Stdout, results)
	case update:
		if opts.message == "" {
			message, err := editorMessage()
			if err != nil {
				log.Fatal(err)
			}
			opts.message = string(message)
		}
		issuePatch := github.Issue{
			Number: opts.issue,
			Body:   string(opts.message),
		}
		issue, err := github.UpdateIssue(opts.token, opts.repo, issuePatch)
		if err != nil {
			log.Fatal(err)
		}
		IssueReport.Execute(os.Stdout, issue)
	default:
		flag.Usage()
		os.Exit(2)
	}
}

// editorMessage calls EDITOR with a tmp file and returns the result.
func editorMessage() ([]byte, error) {

	// check for environment variable
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return nil, errors.New("warning: no EDITOR environment variable")
	}
	// create tmp file & close it
	f, err := ioutil.TempFile("", "github-message")
	if err != nil {
		return nil, err
	}
	fpath := f.Name()
	f.Close()
	// invoke editor
	cmd := exec.Command(editor, fpath)

	// https://stackoverflow.com/questions/12088138/trying-to-launch-an-external-editor-from-within-a-go-program
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	// wait for editor process to finish
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	// read contents of tmp file
	message, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	err = os.Remove(fpath)
	if err != nil {
		return nil, err
	}
	return message, nil
}
