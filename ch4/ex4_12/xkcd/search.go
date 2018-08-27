/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package xkcd

import (
	"fmt"
	"regexp"
)

// Search the comic index for comics having title, number or alt text,
// matching the regular expression re.
func (i *Index) Search(re *regexp.Regexp) []*Comic {
	matches := make([]*Comic, 0, len(i.Comics))
	for _, comic := range i.Comics {
		if re.MatchString(comic.Title) ||
			re.MatchString(comic.Alt) ||
			re.MatchString(fmt.Sprintf("%d", comic.Num)) {
			matches = append(matches, comic)
		}
	}
	return matches
}
