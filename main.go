package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

const (
	exampleData = "./LargeTextData.txt"
	problemSize = 1024 * 1024 // 1MiB
)

var (
	threads int
	size    int
)

func main() {
	//TODO: Let user pick the search algorithm via flag
	flag.IntVar(&threads, "th", 1, "The number of threads to use in a parallel search")
	flag.IntVar(&size, "size", 100, "The size of the problem to solve in MiB")
	flag.Parse()

	fmt.Println("Creating data")
	bytes, answer := generateBufferedData(size * problemSize)

	fmt.Println("Starting search")
	start := time.Now()
	idx, err := parallelSearch(bytes, benny, threads, context.Background())
	//idx, err := benny(bytes)
	t := time.Now()
	elapsed := t.Sub(start)
	if err != nil {
		fmt.Printf("error finding solution using parallel benny approach\n%s\n", err.Error())
	}
	if idx != answer {
		fmt.Printf("Solved the problem incorrectly, answer was %d but we got %d\n", answer, idx)
	}
	fmt.Printf("We found the proper solution %d in %v\n", idx, elapsed)
}
