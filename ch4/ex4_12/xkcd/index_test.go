package xkcd

import (
	"syscall"
	"testing"
)

const TestFileName = "index.test.json"

func TestCreateOrReadIndex(t *testing.T) {
	syscall.Unlink(TestFileName)
	_, err := CreateOrReadIndex(TestFileName)
	if err != nil {
		t.Errorf("CreateOrReadIndex, got %v", err)
	}
	syscall.Unlink(TestFileName)
}

func TestIndex_Save(t *testing.T) {
	syscall.Unlink(TestFileName)
	indexPtr, err := CreateOrReadIndex(TestFileName)
	if err != nil {
		t.Errorf("CreateOrReadIndex failed got %v", err)
	}
	comic := Comic{Title: "hello world"}
	indexPtr.AddComic(comic)
	err = indexPtr.Save(TestFileName)
	if err != nil {
		t.Errorf("Index.Save failed, got %v", err)
	}
	syscall.Unlink(TestFileName)
}

func TestIndex_Update(t *testing.T) {
	syscall.Unlink(TestFileName)
	indexPtr, err := CreateOrReadIndex(TestFileName)
	if err != nil {
		t.Errorf("CreateOrReadIndex failed got %v", err)
	}
	max := 5
	err = indexPtr.Update(max)
	if err != nil {
		t.Errorf("Update failed, got %v", err)
	}
	if len(indexPtr.Comics) != 5 {
		t.Errorf("Expected %d comics, got %d", 5, len(indexPtr.Comics))
	}
	err = indexPtr.Save(TestFileName)
	if err != nil {
		t.Error(err)
	}
	syscall.Unlink(TestFileName)
}
