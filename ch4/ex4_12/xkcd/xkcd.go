/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package xkcd

// APICurrentComic is the URL of latest comic (use to get comic id)
const APICurrentComic = "https://xkcd.com/info.0.json"

// APIFormat is a string template for URL of comid with Id.
const APIFormat = "http://xkcd.com/%d/info.0.json"

// LastComicToken is an alias for "0"
const LastComicToken = 0

// IndexFilename is the name of index file to be created
const IndexFilename = "index.json"

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
