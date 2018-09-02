/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package poster

import (
	"net/url"
)

// DataUrl returns the url for data requests, in the format
// http://www.omdbapi.com/?apikey=[yourkey]&
func DataUrl(apikey string) *url.URL {
	u, _ := url.Parse("http://www.omdbapi.com/")
	q := u.Query()
	q.Set(APIKey.Name, apikey)
	u.RawQuery = q.Encode()
	return u
}

// PosterUrl returns the url for poster api requests, in the format
// http://img.omdbapi.com/?apikey=[yourkey]&
func PosterUrl(apikey string) *url.URL {
	u, _ := url.Parse("http://img.omdbapi.com/")
	q := u.Query()
	q.Set(APIKey.Name, apikey)
	u.RawQuery = q.Encode()
	return u
}
