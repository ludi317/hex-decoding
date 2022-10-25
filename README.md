# hex-decoding

### The quest for the fastest hex decoder written in Go. 
This was the journey to find the all-time fastest hex decoder ever to grace the Go language. 
I benchmarked 9 different implementations of hex decoders in my quest. Some I came up with myself,
with the remaining coming from various sources, including the Go standard library, a Rust crate, and Stack Overflow. 
To be a contender an implementation had to indicate whether the input was valid hex.   
### Benchmarks
And the winner is... a gigantic lookup table!

The fastest approach is to store a 256 x 256 2-D int16 array of precomputed byte values for all possible pairs of hexadecimals. Pairs of hexadecimals simply map to their joint byte value.
For example, table['a']['5'] = 0xa5. Invalid input, meaning any pair with non-hexadecimals, map to -1.

All benchmarks were run on a MacBook Pro (16-inch, 2021) Apple M1 Pro 3220 MHz (10 cores).


```azure
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: hex_decoding
BenchmarkFrom2DInt16-10          	935675250	         1.288 ns/op
BenchmarkFrom1DInt16-10          	886030069	         1.353 ns/op
BenchmarkFrom2Dbyte-10           	867297466	         1.399 ns/op
BenchmarkFromSmallString-10      	841007526	         1.432 ns/op
BenchmarkFromBigString-10        	815906322	         1.483 ns/op
BenchmarkFrom1DByte-10           	813960325	         1.485 ns/op
BenchmarkFrom2SmallStrings-10    	790180687	         1.526 ns/op
BenchmarkFromMath-10             	719872761	         1.671 ns/op
BenchmarkFromBranching-10        	231143634	         5.018 ns/op
PASS
ok  	hex_decoding	12.683s
```

#### Is the 2-D lookup table really the fastest?

Why, yes it is. Running `benchstat` comparing the top 2 methods gives a p-value of 0. This means we can reject the null hypothesis that there's no difference in time between them. 
```azure
$ benchstat 2Dint16.txt 1Dint16.txt 
name     old time/op  new time/op  delta
From-10  1.28ns ± 0%  1.34ns ± 0%  +4.22%  (p=0.000 n=9+9)
```