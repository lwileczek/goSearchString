package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	exampleData = "./LargeTextData.txt"
	problemSize = 1024 * 1024 * 800 // 800MiB
)

func main() {
	if _, err := os.Stat(exampleData); err == nil {
		// path/to/whatever exists
		fmt.Println("File exists you should run benchmarks")

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		generateData(1024*1024*10, exampleData)

	} else {
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		fmt.Println("I don't know what the heck is going on")
		panic(err)

	}
}
