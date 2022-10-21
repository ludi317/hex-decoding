package hex_decoding

import (
	"encoding/hex"
	"math/rand"
	"testing"
)

func TestFrom2D(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hex := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < 2_000_000; j += 2 {
		hi, lo := hex[j], hex[j+1]
		h, ok := From2DInt16(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
}

func TestFrom2Dbyte(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hex := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < 2_000_000; j += 2 {
		hi, lo := hex[j], hex[j+1]
		h, ok := From2Dbyte(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
}

func TestFrom1D(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hex := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < 2_000_000; j += 2 {
		hi, lo := hex[j], hex[j+1]
		h, ok := From1D(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
}

func TestFrom1Dbyte(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hex := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < 2_000_000; j += 2 {
		hi, lo := hex[j], hex[j+1]
		h, ok := From1Dbyte(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
}
func TestFromTableless(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hex := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < 2_000_000; j += 2 {
		hi, lo := hex[j], hex[j+1]
		h, ok := FromTableless(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
}

func TestFromSmallString(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hex := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < 2_000_000; j += 2 {
		hi, lo := hex[j], hex[j+1]
		h, ok := FromSmallString(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
}

func TestFrom2SmallStrings(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hex := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < 2_000_000; j += 2 {
		hi, lo := hex[j], hex[j+1]
		h, ok := From2SmallStrings(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
}

func TestFromBigString(t *testing.T) {
	// construct 2*1MB of hex encoded input, 1MB max (big enough to exceed L1D; small enough to fit in L3)
	bin := make([]byte, 1_000_000)
	rand.Read(bin)
	hex := hex.EncodeToString(bin)

	// decode the hex we just created, and check it matches bin
	for j := 0; j < 2_000_000; j += 2 {
		hi, lo := hex[j], hex[j+1]
		h, ok := FromBigString(hi, lo, true)
		if !ok || bin[j/2] != h {
			t.Fatalf("From(%c, %c) returned (0x%x, %v), expected 0x%x", hi, lo, h, ok, bin[j/2])
		}
	}
}
