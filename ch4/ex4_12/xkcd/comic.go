/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package xkcd

// Comic struct is the serialization container for XKCD JSON files.
type Comic struct {
	Month      string `json:",omitempty"`
	Num        int    `json:",omitempty"`
	Link       string `json:",omitempty"`
	Year       string `json:",omitempty"`
	News       string `json:",omitempty"`
	SafeTitle  string `json:"safe_title,omitempty"`
	Transcript string `json:",omitempty"`
	Alt        string `json:",omitempty"`
	Img        string `json:",omitempty"`
	Title      string `json:",omitempty"`
	Day        string `json:",omitempty"`
}
