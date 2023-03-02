package main

import (
	"context"
	"fmt"
	"math/bits"

	"golang.org/x/sync/errgroup"
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
	return -1, nil //&DidNotFind{FuncName: "naive"}
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
	return -1, nil //&DidNotFind{FuncName: "breakOnDuplicate"}
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
	return -1, nil // &DidNotFind{FuncName: "findWithSlice"}
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
	return -1, nil //&DidNotFind{FuncName: "findWithArray"}
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

	return -1, nil //&DidNotFind{FuncName: "benny"}
}

func davidAPerez(dat []byte) (int, error) {
	var bitIndex byte
	var alreadySet bool
	var state uint32

	position := 0
	for position < len(dat)-14 {
		state = 0 // 0x00000000
		for x := 13; x >= 0; x-- {
			bitIndex = dat[position+x] % 32
			alreadySet = (state & (1 << bitIndex)) != 0
			state |= 1 << bitIndex
			if alreadySet {
				position += x + 1
				break
			}
		}
		if bits.OnesCount32(state) == 14 {
			return position + 14, nil
		}
	}

	return -1, nil //&DidNotFind{FuncName: "davidAPerez"}
}

func parallelSearch(dat []byte, algo searchFunc, thd int, ctx context.Context) (int, error) {
	dataLength := len(dat)
	dataChunkSize := dataLength / thd
	g, ctx := errgroup.WithContext(ctx)
	ch := make(chan int, 1) // buffer of one allows exactly one answer
	for start := 0; start < dataLength-15; start += dataChunkSize {
		start := start
		end := start + dataChunkSize + 25
		if end > dataLength {
			end = dataLength - 1
		}
		if (end - start) <= 14 {
			continue
		}
		g.Go(func() error {
			//do stuff, return errors as needed, or nil
			i, err := algo(dat[start:end])
			if err != nil {
				return err
			}
			if i < 0 {
				return nil
			}
			select {
			case ch <- (i + start):
				// race won, all done
			case <-ctx.Done():
				return ctx.Err()
			default:
				// not an error, just got scooped by another success
			}
			return nil
		})
	}

	go func() {
		g.Wait()
		close(ch) // signal for loop stop
	}()

	// Runs 0 or 1 time, ends when channel closes.
	for val := range ch {
		return val, nil
	}

	// Safe to call Wait twice.
	if err := g.Wait(); err != nil {
		return -1, fmt.Errorf("parallel find: %w", err)
	}

	return -1, &DidNotFind{FuncName: "parallel find"}
}
