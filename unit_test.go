package hex_decoding

import (
	"encoding/hex"
	"math/rand"
	"testing"
)

func TestFrom2DInt16(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < len(hexStr); j += 2 {
		hi, lo := hexStr[j], hexStr[j+1]
		h, ok := From2DInt16(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
	// negative cases
	_, ok := From2DInt16('l', 'r', true)
	if ok {
		t.Fatalf("want not hex, got hex")
	}
}

func TestFrom1DInt16(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < len(hexStr); j += 2 {
		hi, lo := hexStr[j], hexStr[j+1]
		h, ok := From1DInt16(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
	// negative cases
	_, ok := From1DInt16('l', 'r', true)
	if ok {
		t.Fatalf("want not hex, got hex")
	}
}

func TestFrom2Dbyte(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < len(hexStr); j += 2 {
		hi, lo := hexStr[j], hexStr[j+1]
		h, ok := From2Dbyte(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
	// negative cases
	_, ok := From2Dbyte('l', 'r', true)
	if ok {
		t.Fatalf("want not hex, got hex")
	}
}

func TestFrom2SmallArrays(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < len(hexStr); j += 2 {
		hi, lo := hexStr[j], hexStr[j+1]
		h, ok := From2SmallArrays(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}

	// negative cases
	_, ok := From2SmallArrays('l', 'r', true)
	if ok {
		t.Fatalf("want not hex, got hex")
	}
}

func TestFromSmallString(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < len(hexStr); j += 2 {
		hi, lo := hexStr[j], hexStr[j+1]
		h, ok := FromSmallString(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
	// negative cases
	_, ok := FromSmallString('l', 'r', true)
	if ok {
		t.Fatalf("want not hex, got hex")
	}
}

func TestFromBigString(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < len(hexStr); j += 2 {
		hi, lo := hexStr[j], hexStr[j+1]
		h, ok := FromBigString(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
	// negative cases
	_, ok := FromBigString('l', 'r', true)
	if ok {
		t.Fatalf("want not hex, got hex")
	}
}

func TestFrom2SmallStrings(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < len(hexStr); j += 2 {
		hi, lo := hexStr[j], hexStr[j+1]
		h, ok := From2SmallStrings(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
	// negative cases
	_, ok := From2SmallStrings('l', 'r', true)
	if ok {
		t.Fatalf("want not hex, got hex")
	}
}

func TestFrom1Dbyte(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < len(hexStr); j += 2 {
		hi, lo := hexStr[j], hexStr[j+1]
		h, ok := From1Dbyte(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
	// negative cases
	_, ok := From1Dbyte('l', 'r', true)
	if ok {
		t.Fatalf("want not hex, got hex")
	}
}

func TestFromMath(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < len(hexStr); j += 2 {
		hi, lo := hexStr[j], hexStr[j+1]
		h, ok := FromMath(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
	// negative cases
	_, ok := FromMath('l', 'r', true)
	if ok {
		t.Fatalf("want not hex, got hex")
	}
}

func TestFromBranching(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hexStr := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < len(hexStr); j += 2 {
		hi, lo := hexStr[j], hexStr[j+1]
		h, ok := FromBranching(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
	// negative cases
	_, ok := FromBranching('l', 'r', true)
	if ok {
		t.Fatalf("want not hex, got hex")
	}
}
