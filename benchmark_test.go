package hex_decoding

import (
	"encoding/hex"
	"math/rand"
	"testing"
)

var sink byte

const hexInputSize = 1 << 12

func BenchmarkFrom2DInt16(b *testing.B) {
	inputB := make([]byte, hexInputSize/2)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, y := From2DInt16(input[i%hexInputSize], input[(i+1)%hexInputSize], true)
		if y {
			s += x
		}
	}
	sink = s
}

func BenchmarkFrom1DInt16(b *testing.B) {
	inputB := make([]byte, hexInputSize/2)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, y := From1DInt16(input[i%hexInputSize], input[(i+1)%hexInputSize], true)
		if y {
			s += x
		}
	}
	sink = s
}

func BenchmarkFrom2SmallArrays(b *testing.B) {
	inputB := make([]byte, hexInputSize/2)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, y := From2SmallArrays(input[i%hexInputSize], input[(i+1)%hexInputSize], true)
		if y {
			s += x
		}
	}
	sink = s
}

func BenchmarkFrom2Dbyte(b *testing.B) {
	inputB := make([]byte, hexInputSize/2)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, y := From2Dbyte(input[i%hexInputSize], input[(i+1)%hexInputSize], true)
		if y {
			s += x
		}
	}
	sink = s
}

func BenchmarkFromSmallString(b *testing.B) {
	inputB := make([]byte, hexInputSize/2)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, y := FromSmallString(input[i%hexInputSize], input[(i+1)%hexInputSize], true)
		if y {
			s += x
		}
	}
	sink = s
}

func BenchmarkFromBigString(b *testing.B) {
	inputB := make([]byte, hexInputSize/2)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, y := FromBigString(input[i%hexInputSize], input[(i+1)%hexInputSize], true)
		if y {
			s += x
		}
	}
	sink = s
}

func BenchmarkFrom1DByte(b *testing.B) {
	inputB := make([]byte, hexInputSize/2)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, y := From1Dbyte(input[i%hexInputSize], input[(i+1)%hexInputSize], true)
		if y {
			s += x
		}
	}
	sink = s
}

func BenchmarkFrom2SmallStrings(b *testing.B) {
	inputB := make([]byte, hexInputSize/2)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, y := From2SmallStrings(input[i%hexInputSize], input[(i+1)%hexInputSize], true)
		if y {
			s += x
		}
	}
	sink = s
}

func BenchmarkFromMath(b *testing.B) {
	inputB := make([]byte, hexInputSize/2)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, y := FromMath(input[i%hexInputSize], input[(i+1)%hexInputSize], true)
		if y {
			s += x
		}
	}
	sink = s
}

func BenchmarkFromBranching(b *testing.B) {
	inputB := make([]byte, hexInputSize/2)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, y := FromBranching(input[i%hexInputSize], input[(i+1)%hexInputSize], true)
		if y {
			s += x
		}
	}
	sink = s
}
