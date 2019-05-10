package _443_String_Compression

import (
	"testing"
)

func TestStringCompression(t *testing.T) {
	var chars = ([]byte)("abbbbbbbbbbbbbbbbbbbbbbbbbb")
	var len = compress(chars)
	if len != 5 && string(chars) != "a2b2c" {
		t.Fatalf("compress length: %d, compress: %s", len, string(chars))
	}
	t.Logf("compress length: %d, compress: %s", len, string(chars))

	//chars = ([]byte)("aa")
	//len = compress(chars)
	//if len != 6 && string(chars) != "a2b2a2" {
	//	t.Fatalf("compress length: %d, compress: %s", len, string(chars))
	//}
	//t.Logf("compress length: %d, compress: %s", len, string(chars))

	return
}