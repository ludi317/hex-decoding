package hex_decoding

import (
	"encoding/hex"
	"math/rand"
	"testing"
)

var sink byte

const inputSize = 12
const reps = 10000

// generic benchmarker. Swap out its From() func and rerun to compare with benchstat.
func dBenchmarkFrom(b *testing.B) {
	inputB := make([]byte, inputSize)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < reps/inputSize; k++ {
			for j := 0; j < inputSize; j += 2 {
				x, y := From1D(input[j], input[(j+1)], true)
				if y {
					s += x
				}
			}
		}
	}
	sink = s
}

func BenchmarkFrom2DInt(b *testing.B) {
	inputB := make([]byte, inputSize)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < reps/inputSize; k++ {
			for j := 0; j < inputSize; j += 2 {
				x, y := From2DInt(input[j], input[(j+1)], true)
				if y {
					s += x
				}
			}
		}
	}
	sink = s
}

// disabled because FromMath() always assumes input is valid.
// Even with this advantage, it is slower than a lookup table.
func dBenchmarkFromMath(b *testing.B) {
	inputB := make([]byte, inputSize)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < reps/inputSize; k++ {
			for j := 0; j < inputSize; j += 2 {
				x, y := FromMath(input[j], input[(j+1)], true)
				if y {
					s += x
				}
			}
		}
	}
	sink = s
}

func BenchmarkFrom2DInt16(b *testing.B) {
	inputB := make([]byte, inputSize)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < reps/inputSize; k++ {
			for j := 0; j < inputSize; j += 2 {
				x, y := From2DInt16(input[j], input[(j+1)], true)
				if y {
					s += x
				}
			}
		}
	}
	sink = s
}

func BenchmarkFrom1D(b *testing.B) {
	inputB := make([]byte, inputSize)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < reps/inputSize; k++ {
			for j := 0; j < inputSize; j += 2 {
				x, y := From1D(input[j], input[(j+1)], true)
				if y {
					s += x
				}

			}
		}
	}
	sink = s
}

/*
func BenchmarkFromSmallString(b *testing.B) {
	inputB := make([]byte, inputSize)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < reps/inputSize; k++ {
			for j := 0; j < inputSize; j += 2 {
				x, y := FromSmallString(input[j], input[(j+1)], true)
				if y {
					s += x
				}
			}
		}
	}
	sink = s
}

func BenchmarkFrom2Dbyte(b *testing.B) {
	inputB := make([]byte, inputSize)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < reps/inputSize; k++ {
			for j := 0; j < inputSize; j += 2 {
				x, y := From2Dbyte(input[j], input[(j+1)], true)
				if y {
					s += x
				}
			}
		}
	}
	sink = s
}

func BenchmarkFromBigString(b *testing.B) {
	inputB := make([]byte, inputSize)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < reps/inputSize; k++ {
			for j := 0; j < inputSize; j += 2 {
				x, y := FromBigString(input[j], input[(j+1)], true)
				if y {
					s += x
				}
			}
		}
	}
	sink = s
}

func BenchmarkFrom1DByte(b *testing.B) {
	inputB := make([]byte, inputSize)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < reps/inputSize; k++ {

			for j := 0; j < inputSize; j += 2 {
				x, y := From1Dbyte(input[j], input[(j+1)], true)
				if y {
					s += x
				}
			}
		}
	}
	sink = s
}

func BenchmarkFrom2SmallStrings(b *testing.B) {
	inputB := make([]byte, inputSize)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < reps/inputSize; k++ {
			for j := 0; j < inputSize; j += 2 {
				x, y := From2SmallStrings(input[j], input[(j+1)], true)
				if y {
					s += x
				}
			}
		}
	}
	sink = s
}

func BenchmarkFromBranching(b *testing.B) {
	inputB := make([]byte, inputSize)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < reps/inputSize; k++ {
			for j := 0; j < inputSize; j += 2 {
				x, y := FromBranching(input[j], input[(j+1)], true)
				if y {
					s += x
				}
			}
		}
	}
	sink = s
}
*/
