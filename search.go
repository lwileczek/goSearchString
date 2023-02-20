package main

import (
	"fmt"
	"log"
	"math/bits"
	"runtime"
)

type searchFunc func([]byte) (int, error)

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
		if bits.OnesCount32(filter) == 14 {
			return pos + 14, nil
		}
		filter ^= 1 << (first % 32)
	}

	return 0, fmt.Errorf("Never found a proper solution")
}

func davidAPerez(dat []byte) (int, error) {
	var bitIndex byte
	var alreadySet bool

	position := 0
	for position < len(dat)-14 {
		var state uint32
		for x := 13; x >= 0; x-- {
			bitIndex = dat[position+x] % 32
			alreadySet = (state & (1 << bitIndex)) != 0
			state |= 1 << bitIndex
			if alreadySet {
				position += x + 1
				break
			} else if bits.OnesCount32(state) == 14 {
				return position + 14, nil
			}
		}
	}

	return 0, fmt.Errorf("Never found a proper solution")
}

//TODO: Need to add Error Group to this
func parallelFind(dat []byte, algo searchFunc) (int, error) {
	nCPU := runtime.NumCPU()
	dataChunkSize := len(dat) / nCPU
	ch := make(chan int)
	for c := 0; c < nCPU-1; c++ {
		go func(b []byte, startIdx int) {
			i, err := algo(b)
			if err != nil {
				if err.Error() != "Never found a proper solution" {
					log.Println(err)
				}
			}
			ch <- i + startIdx
		}(dat[(c*dataChunkSize):((c+1)*dataChunkSize)], c*dataChunkSize)
	}
	ans := <-ch
	return ans, nil
}
