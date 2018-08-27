/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Index struct is the serialization of the Comic index
type Index struct {
	Comics []*Comic
}

// Save writes the index to storage in json format.
func (i *Index) Save(filename string) error {
	if filename == "" {
		filename = IndexFilename
	}
	bytes, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, bytes, 0644)
	return err
}

// AddComic appends comic to the slice of comics in the index.
func (i *Index) AddComic(comic Comic) {
	i.Comics = append(i.Comics, &comic)
}

// PreviousComic returns the previously fetched comic
func (i *Index) PreviousComic() *Comic {
	if len(i.Comics) == 0 {
		return &Comic{}
	}
	return i.Comics[len(i.Comics)-1]
}

// Update the index by querying xkcd json api for latest. Note, does not Save().
func (i *Index) Update(maxComicId int) error {
	// fetch the latest comic
	latestComic, err := Get(0)
	if err != nil {
		return error(fmt.Errorf("fetch latest comic (id 0) failed: got %v", err))
	}
	// early out if the latestComic is the only one needing to be updated.
	prevComic := i.PreviousComic()
	if prevComic.Num == latestComic.Num-1 {
		i.AddComic(latestComic)
		return nil
	}
	startComicId := prevComic.Num + 1
	for id := startComicId; id < latestComic.Num; id++ {
		if maxComicId > 0 && id > maxComicId {
			break
		}
		if id == 404 {
			// 404 is an easter egg (is not found)
			continue
		}
		log.Printf("get comic %d...", id)
		comic, err := Get(id)
		if err != nil {
			return fmt.Errorf("fetch comic id %d failed: got %v", id, err)
		}
		i.AddComic(comic)
		time.Sleep(time.Second)
	}
	return nil
}

// CreateOrReadIndex creates the index file if it does not exist;
// or returns reads and returns existing index.
func CreateOrReadIndex(filename string) (*Index, error) {
	if filename == "" {
		filename = IndexFilename
	}
	index := Index{
		Comics: make([]*Comic, 0, 10),
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// create index file
		json, err := json.Marshal(index)
		if err != nil {
			return &index, err
		}
		err = ioutil.WriteFile(filename, json, 0644)
		if err != nil {
			return &index, err
		}
		return &index, nil
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return &index, err
	}
	err = json.Unmarshal(bytes, &index)
	if err != nil {
		return &index, err
	}
	return &index, nil
}
