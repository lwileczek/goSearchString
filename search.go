package main

import (
	"fmt"
	"math/bits"
)

//TODO: Read the file in chunks

func naive(dat []byte) (int, error) {
	set := make(UniqueLetters)
	for p := 0; p < len(dat)-14; p++ {
		for w := 0; w < 14; w++ {
			set.Add(dat[p+w])
		}
		if set.Length() == 14 {
			return p + 14, nil
		}
		set.Clear()
	}
	return 0, fmt.Errorf("Never found a proper solution")
}

func checkWindow(s []byte) bool {
	set := make(UniqueLetters)
	for w := 0; w < 14; w++ {
		if set.Has(s[w]) {
			return false
		}
		set.Add(s[w])
	}
	return set.Length() == 14
}

func findWithBreakOnDuplicate(dat []byte) (int, error) {
	maxLength := len(dat) - 14
	for p := 0; p < maxLength; p++ {
		if checkWindow(dat[p : p+14]) {
			return p + 14, nil
		}
	}
	return 0, fmt.Errorf("Never found a proper solution")
}

func checkSliceWindow(s []byte) bool {
	window := make([]byte, 14)
	for w := 0; w < 14; w++ {
		for _, b := range window {
			if b == s[w] {
				return false
			}
		}
		window[w] = s[w]
	}
	return true
}

func findWithSlice(dat []byte) (int, error) {
	for p := 0; p < len(dat)-14; p++ {
		if checkSliceWindow(dat[p : p+14]) {
			return p + 14, nil
		}
	}
	return 0, fmt.Errorf("Never found a proper solution")
}

func checkArrayWindow(s []byte) bool {
	var window [14]byte
	for w := 0; w < 14; w++ {
		for j := 0; j < w; j++ {
			if window[j] == s[w] {
				return false
			}
		}
		window[w] = s[w]
	}
	return true
}

func findWithArray(dat []byte) (int, error) {
	for p := 0; p < len(dat)-14; p++ {
		if checkArrayWindow(dat[p : p+14]) {
			return p + 14, nil
		}
	}
	return 0, fmt.Errorf("Never found a proper solution")
}

func benny(dat []byte) (int, error) {
	var filter uint32
	for k := 0; k < 13; k++ {
		filter ^= 1 << (dat[k] % 32)
	}
	for pos := 0; pos < len(dat)-14; pos++ {
		first := dat[pos]
		filter ^= 1 << (dat[pos+13] % 32)
		if pos < 10 {
			fmt.Printf("%b\n", filter)
		}
		if bits.OnesCount32(filter) == 14 {
			return pos + 14, nil
		}
		filter ^= 1 << (first % 32)
	}

	return 0, fmt.Errorf("Never found a proper solution")
}
