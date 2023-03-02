package main

import (
	"context"
	"testing"
)

var (
	bytes, answer = generateBufferedData(1024 * 1024 * 100) //100MiB
)

func TestNaive(t *testing.T) {
	testBytes, ans := generateBufferedData(4096)
	i, err := naive(testBytes)
	if err != nil {
		t.Error(err)
	}
	if i != ans {
		t.Errorf("Did not properly find the answer. Correct answer is %d but got %d\n", ans, i)
	}
	badBytes := generateNoSolution(1024)
	i, err = naive(badBytes)
	if i != -1 {
		t.Error("Solution did not properly announce that it did not find a solution")
	}
}

func TestHashWithBreak(t *testing.T) {
	testBytes, ans := generateBufferedData(4096)
	i, err := findWithBreakOnDuplicate(testBytes)
	if err != nil {
		t.Error(err)
	}
	if i != ans {
		t.Errorf("Did not properly find the answer. Correct answer is %d but got %d\n", ans, i)
	}
	badBytes := generateNoSolution(1024)
	i, err = findWithBreakOnDuplicate(badBytes)
	if i != -1 {
		t.Error("Solution did not properly announce that it did not find a solution")
	}
}

func TestSliceWindow(t *testing.T) {
	testBytes, ans := generateBufferedData(4096)
	i, err := findWithSlice(testBytes)
	if err != nil {
		t.Error(err)
	}
	if i != ans {
		t.Errorf("Did not properly find the answer. Correct answer is %d but got %d\n", ans, i)
	}
	badBytes := generateNoSolution(1024)
	i, err = findWithSlice(badBytes)
	if i != -1 {
		t.Error("Solution did not properly announce that it did not find a solution")
	}
}

func TestVectorWindow(t *testing.T) {
	testBytes, ans := generateBufferedData(4096)
	i, err := findWithArray(testBytes)
	if err != nil {
		t.Error(err)
	}
	if i != ans {
		t.Errorf("Did not properly find the answer. Correct answer is %d but got %d\n", ans, i)
	}
	badBytes := generateNoSolution(1024)
	i, err = findWithArray(badBytes)
	if i != -1 {
		t.Error("Solution did not properly announce that it did not find a solution")
	}
}

func TestBenny(t *testing.T) {
	testBytes, ans := generateBufferedData(4096)
	i, err := benny(testBytes)
	if err != nil {
		t.Error(err)
	}
	if i != ans {
		t.Errorf("Did not properly find the answer. Correct answer is %d but got %d\n", ans, i)
	}
	badBytes := generateNoSolution(1024)
	i, err = benny(badBytes)
	if i != -1 {
		t.Error("Solution did not properly announce that it did not find a solution")
	}
}

func TestPerez(t *testing.T) {
	testBytes, ans := generateBufferedData(4096)
	i, err := davidAPerez(testBytes)
	if err != nil {
		t.Error(err)
	}
	if i != ans {
		t.Errorf("Did not properly find the answer. Correct answer is %d but got %d\n", ans, i)
	}
	badBytes := generateNoSolution(1024)
	i, err = davidAPerez(badBytes)
	if i != -1 {
		t.Error("Solution did not properly announce that it did not find a solution")
	}
}

func TestParallel(t *testing.T) {
	testBytes, ans := generateBufferedData(4096 * 8)
	i, err := parallelSearch(context.Background(), testBytes, davidAPerez, 4)
	if err != nil {
		t.Error(err)
	}
	if i != ans {
		t.Errorf("Did not properly find the answer. Correct answer is %d but got %d\n", ans, i)
	}
	badBytes := generateNoSolution(4096 * 8)
	i, err = parallelSearch(context.Background(), badBytes, davidAPerez, 4)
	if i != -1 {
		t.Error("Solution did not properly announce that it did not find a solution")
	}
}

func BenchmarkNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := naive(bytes)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkHashWithBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := findWithBreakOnDuplicate(bytes)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSliceWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := findWithSlice(bytes)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkVectorWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := findWithArray(bytes)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkBenny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := benny(bytes)
		if err != nil {
			panic(err)
		}
	}
}
func BenchmarkPerez(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := davidAPerez(bytes)
		if err != nil {
			panic(err)
		}
	}
}
func BenchmarkParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := parallelSearch(context.Background(), bytes, benny, 4)
		if err != nil {
			panic(err)
		}
	}
}
