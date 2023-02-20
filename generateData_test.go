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
	fileSize := 4096
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

func TestGeneratingInMemArray(t *testing.T) {
	size := 1024
	byteArr, n := generateBufferedData(size)
	if len(byteArr) != size {
		t.Error("Did not creat the correct length array")
	}
	u := make(UniqueLetters)
	for _, b := range byteArr {
		u.Add(b)
	}
	if u.Length() < 14 {
		t.Error("Not possible to find solution")
	}
	u.Clear()
	for i := 14; i > 0; i-- {
		u.Add(byteArr[n-i])
	}
	if u.Length() != 14 {
		t.Error("Solution is not at the expected location")
	}
}

func TestGeneratingNoSolution(t *testing.T) {
	size := 1024
	byteArr := generateNoSolution(size)
	if len(byteArr) != size {
		t.Error("Did not creat the correct length array")
	}
	u := make(UniqueLetters)
	for _, b := range byteArr {
		u.Add(b)
	}
	if u.Length() > 13 {
		t.Error("Possible solution created!")
	}
}
