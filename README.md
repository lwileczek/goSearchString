# Go Search String
Searching through strings quickly.

This repository was inspired by two others:

 - https://github.com/benhoyt/countwords/blob/master/optimized.go
 - https://github.com/ThePrimeagen/aoc/blob/2022/src/bin/day6\_2.rs

I read a blog post from Ben Hoyt where we talked about I/O was no longer the bottle neck and he talked a little about fast text search 
[See here](https://benhoyt.com/writings/io-is-no-longer-the-bottleneck/). 
I then saw [this YouTube](https://www.youtube.com/watch?v=U16RnpV48KQ) From the Primeagen.
These made me want to work through writing the solutions in Go since I love go and we can easily benchmark
each function with Go's standard utilities.

## Data
To generate data I wrote a little function that will make a rather large 800MiB file where we 
guarentee a solution at the end if it does not randomly happen before.

```bash
$ go run main.go generate_data.go
```

## Benchmarks

```
$ go test ./... -bench=. -benchmem
```

### Results
Your results will vary a little depending on your CPU but the orders of magnitude should not
```base
# Created a 100 MiB array
go test ./... -bench=.
2023/02/19 22:52:17 Answer Index: 1835022
goos: darwin
goarch: amd64
pkg: github.com/lwileczek/goSearchString
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
BenchmarkNaive-8           	       1	74445413180 ns/op
BenchmarkHashWithBreak-8   	       1	12842798109 ns/op
BenchmarkSliceWindow-8     	       1	 2405292994 ns/op
BenchmarkVectorWindow-8    	       1	 1135809643 ns/op
BenchmarkBenny-8           	      14	   81657561 ns/op
BenchmarkParez-8           	      10	  109475645 ns/op
BenchmarkParallel-8        	      91	   21385322 ns/op
PASS
ok  	github.com/lwileczek/goSearchString	99.112s
```
|Algorithm | x Faster | % Faster |
|:---|---:|---|
|BenchmarkNaive | 0 | 0 |
|BenchmarkHashWithBreak | 4.8 | 480% |
|BenchmarkSliceWindow |  30.0 |	2995% |
|BenchmarkVectorWindow |  64.5 |	6454% |
|BenchmarkBenny |  910.7 |	91,068% |
|BenchmarkParez |  679.0 | 	67,902% |
|BenchmarkParallel (8 cpu) |  3,480.1 | 348, 015% |

I'm not sure if the degradation in the David A Perez algorithm is because I did not program it correctly
or a difference in Go vs. Rust.
In the video with the Primeagen, he was showing off what the compiler was doing with reverse iterators
where I just did a reverse for loop.
