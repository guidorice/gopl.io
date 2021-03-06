/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package dedup

type stringSlice []string

// Dedup eliminates duplicates in the receiver. https://blog.golang.org/slices
// says It is idiomatic to use a pointer receiver for a method that modifies a
// slice.
func (s *stringSlice) Dedup() {
	n := 0
	for _, current := range *s {
		if n > 0 && current == (*s)[n-1] {
			// adjacent duplicate, skip it
			continue
		}
		(*s)[n] = current
		n++
	}
	*s = (*s)[:n]
}

// DedupInPlace eliminates (in-place) duplicates in a slice of strings. Warning:
// original array is modified!
func DedupInPlace(s *[]string) {
	n := 0
	for _, current := range *s {
		if n > 0 && current == (*s)[n-1] {
			// adjacent duplicate, skip it
			continue
		}
		(*s)[n] = current
		n++
	}
	*s = (*s)[:n]
}

// Dedup eliminates duplicates in a slice of strings.
func Dedup(s []string) []string {
	var res []string
	for _, current := range s {
		if len(res) > 0 && current == res[len(res)-1] {
			// adjacent duplicate, skip it
			continue
		}
		res = append(res, current)
	}
	return res
}
