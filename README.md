# hex-decoding

## The quest for the fastest hex decoder written in Go. 
## Benchmarks
All benchmarks ran on a MacBook Pro (16-inch, 2021)
Apple M1 Pro 3220 MHz (10 cores).
### When the input isn't too big
The fastest approach is to store a 256 x 256 2-D int array of precomputed byte values for all possible pairs of hexadecimals. Pairs of hexadecimals simply map to their joint byte value.
For example, table['a']['5'] = 0xa5. Invalid input, meaning any pair with non-hexadecimals, map to -1.

```azure
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: hex_decoding
BenchmarkFrom2DInt_Normal-10            	11275908	       104.1 ns/op
BenchmarkFrom2DInt16_Normal-10          	11471845	       103.9 ns/op
BenchmarkFrom1D_Normal-10               	11453546	       104.4 ns/op
BenchmarkFromSmallString_Normal-10      	10260147	       116.8 ns/op
BenchmarkFrom2Dbyte_Normal-10           	 9968356	       120.3 ns/op
BenchmarkFromBigString_Normal-10        	 9617901	       125.0 ns/op
BenchmarkFrom1DByte_Normal-10           	 9577893	       125.2 ns/op
BenchmarkFrom2SmallStrings_Normal-10    	 7255498	       164.7 ns/op
BenchmarkFromTableless_Normal-10        	 3194462	       381.2 ns/op


```
### When the input is too big
The fastest approach is similar to before, but with a slight tweak. The 2-D precomputed table gets smaller, by changing its type from int to byte. 
-1 can no longer indicate invalid input since that would overflow a byte. Instead, non-hexadecimal maps to 0, as does the valid input pair ('0', '0'). If the table returns 0, an extra check on the input is required to check for validity.
Under intense memory pressure, the extra runtime check to disambiguate the 0 case is worth shrinking the size of the lookup table.  

```azure
BenchmarkFrom2Dbyte-10                  	    1413	    847448 ns/op
BenchmarkFromBigString-10               	    1340	    886595 ns/op
BenchmarkFrom1Dbyte-10                  	    1347	    886813 ns/op
BenchmarkFrom2D-10                      	    1284	    933001 ns/op
BenchmarkFrom1D-10                      	    1280	    934242 ns/op
BenchmarkFromSmallString-10             	    1230	    973432 ns/op
BenchmarkFrom2SmallStrings-10           	     961	   1246867 ns/op
BenchmarkFromTableless-10               	     180	   6660506 ns/op
```