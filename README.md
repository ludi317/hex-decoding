# hex-decoding

### The quest for the fastest hex decoder written in Go. 
This was the journey to find the all-time fastest hex decoder ever to grace the Go language. 
I benchmarked 10 different implementations of hex decoders in my quest. Some I came up with myself,
with the remaining coming from various sources, including the Go standard library, a Rust crate, a coworker, and Stack Overflow. 
To be a contender an implementation had to indicate whether the input was valid hex.   
### Benchmarks
And the winner is... a gigantic lookup table!

The fastest approach is to store a 256 x 256 2D int16 array of precomputed byte values for all possible pairs of hexadecimals. 
Pairs of hexadecimals simply map to their joint byte value.
For example, table['a']['5'] = 0xa5. Invalid input, meaning any pair with non-hexadecimals, map to -1.

Benchmarks were run on a MacBook Pro (16-inch, 2021) Apple M1 Pro 3220 MHz (10 cores).


```azure
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: hex_decoding
BenchmarkFrom2DInt16-10          	925948252	         1.282 ns/op
BenchmarkFrom1DInt16-10          	894526339	         1.340 ns/op
BenchmarkFrom2SmallArrays-10     	878882383	         1.365 ns/op
BenchmarkFrom2Dbyte-10           	869709645	         1.398 ns/op
BenchmarkFromSmallString-10      	845103649	         1.418 ns/op
BenchmarkFromBigString-10        	827684672	         1.491 ns/op
BenchmarkFrom2SmallStrings-10    	814923741	         1.465 ns/op
BenchmarkFrom1DByte-10           	829009581	         1.446 ns/op
BenchmarkFromMath-10             	626569141	         1.918 ns/op
BenchmarkFromBranching-10        	245880645	         4.715 ns/op
```
### FAQ 
#### Is the 2D lookup table really the fastest?

It is on these benchmarks. Running `benchstat` comparing the top 2 methods gives a p-value of 0. This means we can reject the null hypothesis that there's no difference in time between them. 
```azure
$ benchstat 2Dint16.txt 1Dint16.txt 
name     old time/op  new time/op  delta
From-10  1.28ns ± 0%  1.34ns ± 0%  +4.22%  (p=0.000 n=9+9)
```

#### Should I use the 2D table method in my production code to decode hex?
Probably not. I would personally go with From2SmallArrays.
It uses 2 small lookup tables, each with length 256. The faster methods have lookup tables with length 256*256, which I think uses too much memory for the sake of speed.

#### What about when I have 2MB of hex to decode, and I saturate my L1, L2, and L3 caches?
The fastest method is still a 2D lookup table, but no longer one that stores int16s. 
The 2D precomputed table gets smaller, by changing its type from int16 to byte. 
-1 can't indicate invalid input anymore since that would underflow a byte. 
Instead, non-hexadecimal maps to 0, as does the valid input pair ('0', '0'). 
If the table returns 0, an extra check on the input is required to check for validity. 
Under intense memory pressure, the extra runtime check to disambiguate the 0 case is worth shrinking the size of the lookup table.

#### Why is a 2D table faster than a flattened version of itself, as a 1D single array? Shouldn't they be the same?
An examination of the assembly code reveals that the 1-D lookup generates an extra opcode as compared to the 2D. Pure array indexing is better optimized by the compiler.