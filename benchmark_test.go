package hex_decoding

import (
	"encoding/hex"
	"math/rand"
	"testing"
)

var sink byte

func BenchmarkFrom2DInt_Normal(b *testing.B) {
	inputB := make([]byte, 256)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 256; j += 2 {
			x, y := From2DInt(input[j], input[(j+1)], true)
			if y {
				s += x
			}
		}
	}
	sink = s
}

func BenchmarkFrom2DInt16_Normal(b *testing.B) {
	inputB := make([]byte, 256)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 256; j += 2 {
			x, y := From2DInt16(input[j], input[(j+1)], true)
			if y {
				s += x
			}
		}
	}
	sink = s
}

func BenchmarkFrom1D_Normal(b *testing.B) {
	inputB := make([]byte, 256)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		for j := 0; j < 256; j += 2 {
			x, y := From1D(input[j], input[(j+1)], true)
			if y {
				s += x
			}
		}
	}
	sink = s
}

func BenchmarkFromSmallString_Normal(b *testing.B) {
	inputB := make([]byte, 256)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		for j := 0; j < 256; j += 2 {
			x, y := FromSmallString(input[j], input[(j+1)], true)
			if y {
				s += x
			}
		}
	}
	sink = s
}

func BenchmarkFrom2Dbyte_Normal(b *testing.B) {
	inputB := make([]byte, 256)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 256; j += 2 {
			x, y := From2Dbyte(input[j], input[(j+1)], true)
			if y {
				s += x
			}
		}
	}
	sink = s
}

func BenchmarkFromBigString_Normal(b *testing.B) {
	inputB := make([]byte, 256)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 256; j += 2 {
			x, y := FromBigString(input[j], input[(j+1)], true)

			if y {
				s += x
			}
		}
	}
	sink = s
}

func BenchmarkFrom1DByte_Normal(b *testing.B) {
	inputB := make([]byte, 256)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 256; j += 2 {
			x, y := From1Dbyte(input[j], input[(j+1)], true)
			if y {
				s += x
			}
		}
	}
	sink = s
}

func BenchmarkFrom2SmallStrings_Normal(b *testing.B) {
	inputB := make([]byte, 256)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		for j := 0; j < 256; j += 2 {
			x, y := From2SmallStrings(input[j], input[(j+1)], true)
			if y {
				s += x
			}
		}
	}
	sink = s
}

func BenchmarkFromTableless_Normal(b *testing.B) {
	inputB := make([]byte, 256)
	rand.Read(inputB)
	input := hex.EncodeToString(inputB)

	var s byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 256; j += 2 {
			x, y := FromTableless(input[j], input[(j+1)], true)
			if y {
				s += x
			}
		}
	}
	sink = s
}

func BenchmarkFrom2Dbyte(b *testing.B) {
	b.StopTimer()
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	sum := _sum
	ok := true
	var h byte

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(hexStr); j += 2 {
			hi, lo := hexStr[j], hexStr[j+1]
			h, ok = From2Dbyte(hi, lo, ok)
			sum += h
		}
	}
	_ok = ok
	_sum = sum
}

func BenchmarkFromBigString(b *testing.B) {
	b.StopTimer()
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	sum := _sum
	ok := true
	var h byte

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(hexStr); j += 2 {
			hi, lo := hexStr[j], hexStr[j+1]
			h, ok = FromBigString(hi, lo, ok)
			sum += h
		}
	}
	_ok = ok
	_sum = sum
}

func BenchmarkFrom1Dbyte(b *testing.B) {
	b.StopTimer()
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	sum := _sum
	ok := true
	var h byte

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(hexStr); j += 2 {
			hi, lo := hexStr[j], hexStr[j+1]
			h, ok = From1Dbyte(hi, lo, ok)
			sum += h
		}
	}
	_ok = ok
	_sum = sum
}

func BenchmarkFrom2D(b *testing.B) {
	b.StopTimer()
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	sum := _sum
	ok := true
	var h byte

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(hexStr); j += 2 {
			hi, lo := hexStr[j], hexStr[j+1]
			h, ok = From2DInt16(hi, lo, ok)
			sum += h
		}
	}
	_ok = ok
	_sum = sum
}

var (
	_ok  bool
	_sum byte
)

func BenchmarkFrom1D(b *testing.B) {
	b.StopTimer()
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	sum := _sum
	ok := true
	var h byte

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(hexStr); j += 2 {
			hi, lo := hexStr[j], hexStr[j+1]
			h, ok = From1D(hi, lo, ok)
			sum += h
		}
	}
	_ok = ok
	_sum = sum
}

func BenchmarkFromSmallString(b *testing.B) {
	b.StopTimer()
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	sum := _sum
	ok := true
	var h byte

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(hexStr); j += 2 {
			hi, lo := hexStr[j], hexStr[j+1]
			h, ok = FromSmallString(hi, lo, ok)
			sum += h
		}
	}
	_ok = ok
	_sum = sum
}

func BenchmarkFrom2SmallStrings(b *testing.B) {
	b.StopTimer()
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	sum := _sum
	ok := true
	var h byte

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(hexStr); j += 2 {
			hi, lo := hexStr[j], hexStr[j+1]
			h, ok = From2SmallStrings(hi, lo, ok)
			sum += h
		}
	}
	_ok = ok
	_sum = sum
}

func BenchmarkFromTableless(b *testing.B) {
	b.StopTimer()
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	sum := _sum
	ok := true
	var h byte

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(hexStr); j += 2 {
			hi, lo := hexStr[j], hexStr[j+1]
			h, ok = FromTableless(hi, lo, ok)
			sum += h
		}
	}
	_ok = ok
	_sum = sum
}
