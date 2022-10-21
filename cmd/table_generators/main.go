package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"tst115"
)

func main() {
	//generateLookupTable(1, 0)
	//generateLookupTable(2, -1)
	generateBigString()
}

func generateLookupTable(dimensions int, invalidValue int) {

	if dimensions != 1 && dimensions != 2 {
		panic("dimensions must be 1 or 2")
	}

	bigHexTable := [256][256]int{}
	bigHexArray := [256 * 256]int{}

	for i := range bigHexTable {
		for j := range bigHexTable[i] {
			bigHexTable[i][j] = invalidValue
			bigHexArray[i<<8+j] = invalidValue
		}
	}
	hexadecimals := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F'}
	for _, hi := range hexadecimals {
		for _, lo := range hexadecimals {
			v, _ := hex.DecodeString(string([]byte{hi, lo}))
			bigHexTable[hi][lo] = int(v[0])
			bigHexArray[(uint16(hi)<<8 + uint16(lo))] = int(v[0])
		}
	}

	tableType := "uint16"
	switch {
	case invalidValue < 0:
		tableType = "int16"
		if int(int16(invalidValue)) != invalidValue {
			panic(fmt.Sprintf("%d overflows an int16", invalidValue))
		}
	case invalidValue < 0x10:
		tableType = "byte"
	default:
		if int(uint16(invalidValue)) != invalidValue {
			panic(fmt.Sprintf("%d overflows an uint16", invalidValue))
		}
	}

	var outBytes []byte

	if dimensions == 1 {
		outBytes, _ = json.Marshal(bigHexArray)
		fmt.Printf("var Big%sArray = [256*256]%s", strings.Title(tableType), tableType)
	} else if dimensions == 2 {
		outBytes, _ = json.Marshal(bigHexTable)
		fmt.Printf("var Big%sTable = [256][256]%s", strings.Title(tableType), tableType)
	}

	s := string(outBytes)
	s = strings.ReplaceAll(s, " ", ", ")
	s = strings.ReplaceAll(s, "[", "{")
	s = strings.ReplaceAll(s, "]", "}")
	fmt.Println(s)
}

func generateBigString() {
	var BigString = string(hex_decoding.BigByteArray[:])
	s := fmt.Sprintf("% x", []byte(BigString))
	s = strings.ReplaceAll(s, " ", "\\x")
	fmt.Printf("const BigString = \"\\x%s\"\n", s)
}
