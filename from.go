package hex_decoding

func From2DInt16(hi, lo byte, ok bool) (byte, bool) {
	v := BigInt16Table[hi][lo]
	return byte(v), ok && v >= 0
}

func From1DInt16(hi, lo byte, ok bool) (byte, bool) {
	v := BigInt16Array[(uint16(hi)<<8 | uint16(lo))]
	return byte(v), ok && v >= 0
}

func From2SmallArrays(hi, lo byte, ok bool) (byte, bool) {
	hiInt16 := reverseHexTableHiInt16[hi]
	loInt16 := reverseHexTableLoInt16[lo]
	return byte(hiInt16 | loInt16), ok && hiInt16|loInt16 >= 0
}

func From2Dbyte(hi, lo byte, ok bool) (byte, bool) {
	v := BigByteTable[hi][lo]
	if v != 0 {
		return v, ok
	}
	if hi == '0' && lo == '0' {
		return 0, ok
	}
	return 0, false
}

func FromSmallString(hi, lo byte, ok bool) (byte, bool) {
	hi = reverseHexTableLo[hi]
	lo = reverseHexTableLo[lo]
	return (hi << 4) | lo, ok && int8(hi|lo) >= 0
}

func FromBigString(hi, lo byte, ok bool) (byte, bool) {
	v := BigString[(uint16(hi)<<8 | uint16(lo))]
	if v != 0 {
		return v, ok
	}
	if hi == '0' && lo == '0' {
		return 0, ok
	}
	return 0, false
}

func From2SmallStrings(hi, lo byte, ok bool) (byte, bool) {
	hiNew := reverseHexTableHi[hi]
	loNew := reverseHexTableLo[lo]
	return hiNew | loNew, ok && isHexadecimal[hi] && isHexadecimal[lo]
}

func From1Dbyte(hi, lo byte, ok bool) (byte, bool) {
	v := BigByteArray[(uint16(hi)<<8 | uint16(lo))]
	if v != 0 {
		return v, ok
	}
	if hi == '0' && lo == '0' {
		return 0, ok
	}
	return 0, false
}

func FromMath(hi, lo byte, ok bool) (byte, bool) {
	var ok1, ok2 bool
	hi, ok1 = fromNibbleMath(hi)
	lo, ok2 = fromNibbleMath(lo)
	return (hi << 4) | lo, ok && ok1 && ok2
}

func fromNibbleMath(d byte) (byte, bool) {
	// Here's how it works.
	// ascii numbers: 0x30, ..., 0x39
	// ascii letters: 0x41, ..., 0x46
	//              : 0x61, ..., 0x66
	// (d >> 6) = 1 for ascii letters, and 0 for ascii numbers
	// Get lower nibble of ascii number. Done.
	// Get lower nibble for ascii letters. Add 8 + 1.
	isLetter := d >> 6
	return (d&0xf | isLetter<<3) + isLetter, isHexadecimal[d]
}

func FromBranching(hi, lo byte, ok bool) (byte, bool) {
	var ok1, ok2 bool
	hi, ok1 = fromNibbleBranching(hi)
	lo, ok2 = fromNibbleBranching(lo)
	return (hi << 4) | lo, ok && ok1 && ok2
}

func fromNibbleBranching(r byte) (byte, bool) {
	if r <= '9' {
		return r - '0', r >= '0'
	}
	if r >= 'a' {
		return r - 'a' + 10, r <= 'f'
	}
	return r - 'A' + 10, 'A' <= r && r <= 'F'
}
