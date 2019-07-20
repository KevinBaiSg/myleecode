package _14_Longest_Common_Prefix

import (
"testing"
)

func TestStringCompression(t *testing.T) {
	var (
		strs []string
		prefix string
	)

	//prefix = longestCommonPrefix(nil)
	//if prefix != "" {
	//	t.Fatalf("prefix: %s", prefix)
	//}
	//t.Logf("prefix is empty")
	//
	//prefix = longestCommonPrefix(strs)
	//if prefix != "" {
	//	t.Fatalf("prefix: %s", prefix)
	//}
	//t.Logf("prefix is empty")

	strs = append(strs, "flower")
	strs = append(strs, "flow")
	strs = append(strs, "flight")
	prefix = longestCommonPrefix(strs)
	if prefix != "fl" {
		t.Fatalf("prefix: %s", prefix)
	}
	t.Logf("prefix: %s", prefix)

	return
}