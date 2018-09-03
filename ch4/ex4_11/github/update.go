/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// FIXME patch http call is not working

// UpdateIssue updates an issue using github api and the token, repo, and
// existing issue struct.
func UpdateIssue(token string, repo string, issue Issue) (Issue, error) {
	url := fmt.Sprintf(
		"%srepos/%s/issues/%d",
		APIURL,
		repo,
		issue.Number,
	)
	b, err := json.Marshal(issue)
	if err != nil {
		return Issue{}, err
	}
	body := bytes.NewBuffer(b)
	client := &http.Client{}
	req, _ := http.NewRequest("PATCH", string(url), body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+string(token))
	resp, err := client.Do(req)
	if err != nil {
		return Issue{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		limitedReader := &io.LimitedReader{R: resp.Body, N: KiB}
		msg, _ := ioutil.ReadAll(limitedReader)
		return Issue{}, fmt.Errorf("UpdateIssue: failed with status %s, msg %s", resp.Status, msg)
	}
	limitedReader := &io.LimitedReader{R: resp.Body, N: MiB}
	data, err := ioutil.ReadAll(limitedReader)
	if err != nil {
		return Issue{}, err
	}
	updatedIssue := Issue{}
	err = json.Unmarshal(data, &updatedIssue)
	if err != nil {
		return Issue{}, err
	}
	return updatedIssue, nil
}
