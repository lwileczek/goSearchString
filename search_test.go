package main

import (
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
}

func BenchmarkVectorWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := findWithArray(bytes)
		if err != nil {
			panic(err)
		}
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

func BenchmarkSliceWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := findWithSlice(bytes)
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

func BenchmarkBenny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := benny(bytes)
		if err != nil {
			panic(err)
		}
	}
}
