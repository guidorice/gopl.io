/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package poster

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	//GiB
	//TiB
	//PiB
	//EiB
	//ZiB
	//YiB
)

// MovieRequest returns a Movie from the given params (i.e. Id or Title search)
// it returns a Movie struct, or an error.
func MovieRequest(apiKey string, params Params) (Movie, error) {
	url := DataUrl(apiKey)
	q := url.Query()
	for param, value := range params {
		q.Set(param.Name, value)
	}
	url.RawQuery = q.Encode()
	resp, err := http.Get(url.String()) // TODO: add timeouts to http client; dont use defaults
	if err != nil {
		return Movie{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		limitedReader := &io.LimitedReader{R: resp.Body, N: KiB}
		msg, _ := ioutil.ReadAll(limitedReader)
		err := fmt.Errorf(
			"omdb movie fetch failed: status %s, msg %s",
			resp.Status,
			msg,
		)
		return Movie{}, err
	}
	limitedReader := &io.LimitedReader{R: resp.Body, N: MiB}
	data, err := ioutil.ReadAll(limitedReader)
	movie := Movie{}
	err = json.Unmarshal(data, &movie)
	if err != nil {
		return Movie{}, err
	}
	return movie, nil
}

// MovieRequest returns a Movie from the given params (i.e. Id or Title search)
// it returns a SearchResult struct, or an error.
func SearchRequest(apiKey string, params Params) (SearchResult, error) {
	url := DataUrl(apiKey)
	q := url.Query()
	for param, value := range params {
		q.Set(param.Name, value)
	}
	url.RawQuery = q.Encode()
	resp, err := http.Get(url.String()) // TODO: add timeouts to http client; dont use defaults
	if err != nil {
		return SearchResult{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		limitedReader := &io.LimitedReader{R: resp.Body, N: KiB}
		msg, _ := ioutil.ReadAll(limitedReader)
		err := fmt.Errorf(
			"omdb movie search failed: status %s, msg %s",
			resp.Status,
			msg,
		)
		return SearchResult{}, err
	}
	limitedReader := &io.LimitedReader{R: resp.Body, N: MiB}
	data, err := ioutil.ReadAll(limitedReader)
	result := SearchResult{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return SearchResult{}, err
	}
	return result, nil
}

// PosterRequest returns a JPEG movie poster (image.Image) for the omdb movie id.
func PosterRequest(apiKey string, id string) (image.Image, error) {
	url := PosterUrl(apiKey)
	q := url.Query()
	q.Set(Id.Name, id)
	url.RawQuery = q.Encode()
	resp, err := http.Get(url.String()) // TODO: add timeouts to http client; dont use defaults
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		limitedReader := &io.LimitedReader{R: resp.Body, N: KiB}
		msg, _ := ioutil.ReadAll(limitedReader)
		err := fmt.Errorf(
			"omdb poster search failed: status %s, msg %s",
			resp.Status,
			msg,
		)
		return nil, err
	}
	limitedReader := &io.LimitedReader{R: resp.Body, N: MiB}
	data, err := ioutil.ReadAll(limitedReader)
	reader := bytes.NewReader(data)
	img, _, err := image.Decode(reader)
	return img, err
}
