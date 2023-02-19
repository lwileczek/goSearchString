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
...
