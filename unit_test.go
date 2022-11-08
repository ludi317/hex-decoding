package hex_decoding

import (
	"encoding/hex"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestFrom(t *testing.T) {
	for _, from := range []func(hi, lo byte, ok bool) (byte, bool){
		From2DInt16,
		From1DInt16,
		From2Dbyte,
		From2SmallArrays,
		FromSmallString,
		FromBigString,
		From2SmallStrings,
		From1Dbyte,
		FromMath,
		FromBranching,
	} {
		from := from // capture the range variable
		t.Run(getFunctionName(from), func(t *testing.T) {
			t.Parallel()
			for hi := 0; hi < 1<<8; hi++ {
				for lo := 0; lo < 1<<8; lo++ {
					gotVal, ok := from(byte(hi), byte(lo), true)
					// compare our result with the std library's
					wantVal, err := hex.DecodeString(string([]byte{byte(hi), byte(lo)}))
					if ok {
						if err != nil {
							t.Fatalf("want: invalid hex, got: valid hex for hi: %v, lo: %v", hi, lo)
						}
						if gotVal != wantVal[0] {
							t.Fatalf("want: %v, got: %v for hi: %v, lo: %v", wantVal, gotVal, hi, lo)
						}
					} else { // !ok
						if err == nil {
							t.Fatalf("want: valid hex, got: invalid hex for hi: %v, lo: %v", hi, lo)
						}
					}
				}
			}
		})
	}

}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

const signature = " þÿ £ÜÐÎ "

func TestHiddenMessage(t *testing.T) {
	if !strings.Contains(reverseHexTableLo, signature) {
		t.Fatalf("table missing watermark")
	}
	if !strings.Contains(reverseHexTableHi, signature) {
		t.Fatalf("table missing watermark")
	}
}
