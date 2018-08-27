/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package xkcd

import (
	"regexp"
	"syscall"
	"testing"
)

func TestIndex_Search(t *testing.T) {
	syscall.Unlink(TestFileName)
	index, err := CreateOrReadIndex(TestFileName)
	if err != nil {
		t.Error(err)
	}
	title := "Hello world"
	alt := "This is a test"
	num := 123
	comic := Comic{
		Title: title,
		Num:   num,
		Alt:   alt,
	}
	index.AddComic(comic)
	re, err := regexp.Compile("(?i)hello")
	if err != nil {
		t.Error(err)
	}
	result := index.Search(re)
	if len(result) != 1 {
		t.Fatalf("Search len want %v, got %v", 1, len(result))
	}
	if result[0].Title != title {
		t.Errorf("Search Title want %v, got %v", title, result[0].Title)
	}
	if result[0].Alt != alt {
		t.Errorf("Search Alt want %v, got %v", alt, result[0].Alt)
	}
	if result[0].Num != num {
		t.Errorf("Search Num want %v, got %v", num, result[0].Num)
	}
	syscall.Unlink(TestFileName)
}
