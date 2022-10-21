package hex_decoding

func From1D(hi, lo byte, ok bool) (byte, bool) {
	v := BigInt16Array[(uint16(hi)<<8 + uint16(lo))]
	return byte(v), ok && v >= 0
}

func From1Dbyte(hi, lo byte, ok bool) (byte, bool) {
	v := BigByteArray[(uint16(hi)<<8 + uint16(lo))]
	if v != 0 {
		return v, ok
	}
	if hi == '0' && lo == '0' {
		return 0, ok
	}
	return 0, false
}

func From2DInt16(hi, lo byte, ok bool) (byte, bool) {
	v := BigInt16Table[hi][lo]
	return byte(v), ok && v >= 0
}

func From2DInt(hi, lo byte, ok bool) (byte, bool) {
	v := BigIntTable[hi][lo]
	return byte(v), ok && v >= 0
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

func FromTableless(hi, lo byte, ok bool) (byte, bool) {
	var ok1, ok2 bool
	hi, ok1 = fromNibble(hi)
	lo, ok2 = fromNibble(lo)
	return (hi << 4) | lo, ok && ok1 && ok2
}

func fromNibble(r byte) (byte, bool) {
	if r <= '9' {
		return r - '0', r >= '0'
	}
	if r >= 'a' {
		return r - 'a' + 10, r <= 'f'
	}
	return r - 'A' + 10, 'A' <= r && r <= 'F'
}

func FromSmallString(hi, lo byte, ok bool) (byte, bool) {
	hi = reverseHexTableLo[hi]
	lo = reverseHexTableLo[lo]
	return (hi << 4) | lo, ok && int8(hi|lo) >= 0
}

func From2SmallStrings(hi, lo byte, ok bool) (byte, bool) {
	hi = reverseHexTableHi[hi]
	lo = reverseHexTableLo[lo]
	return hi | lo, ok && hi <= 0xf0 && lo <= 0x0f
}

func FromBigString(hi, lo byte, ok bool) (byte, bool) {
	v := BigString[(uint16(hi)<<8 + uint16(lo))]
	if v != 0 {
		return v, ok
	}
	if hi == '0' && lo == '0' {
		return 0, ok
	}
	return 0, false
}
