package main

import (
	"errors"
	"log"
	"os"
	"testing"
)

const (
	testFile = "testfile.txt"
)

func TestGenerateData(t *testing.T) {
	fileSize := 1024 * 1024 * 2
	if _, err := os.Stat(testFile); errors.Is(err, os.ErrNotExist) {
		generateData(fileSize, testFile)
	} else {
		t.Log("File already exists or some other error occurred")
		t.Error(err)
		return
	}
	fileInfo, err := os.Stat(testFile)
	if err != nil {
		t.Error(err)
	}
	if fileInfo.Size() != int64(fileSize) {
		t.Error("Did not generate a file of the correct size")
	}
	//TODO: Check if characters are all lowercase english letters?
	//TODO: Check the end of the file is a valid solution
	if err := os.Remove(testFile); err != nil {
		log.Fatal(err)
	}
}
