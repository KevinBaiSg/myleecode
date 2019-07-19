package solution1

import (
	"testing"
)

func TestZigZagConversion(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	cases := []struct {
		s 		string
		numRows int
		want 	string
	}{
		{"", 10, ""},
		{"LEETCODEISHIRING", 0, ""},
		{"LEETCODEISHIRING", 1, "LEETCODEISHIRING"},
		{"LEETCODEISHIRING", 3, "LCIRETOESIIGEDHN"},
		{"LEETCODEISHIRING", 4, "LDREOEIIECIHNTSG"},
	}
	for _, c := range cases {
		zigZag := convert(c.s, c.numRows)
		if zigZag != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.s, zigZag, c.want)
		}
	}
	return
}

func BenchmarkZigZagConversion(b *testing.B)  {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convert("LEETCODEISHIRING", 4)
	}
	return
}
