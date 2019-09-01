package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	/*
		0 -> 000
		1 -> 001
		2 -> 010
		3 -> 011
		4 -> 100
		5 -> 101
		6 -> 110
		7 -> 111
	*/
	cases := []struct {
		in 		uint32
		want 	uint32
	}{
		{0, 0},
		{01, 020000000000},
		{05, 024000000000},
		{0345, 024700000000},
	}
	for _, c := range cases {
		result := reverseBits(c.in)
		if result != c.want {
			t.Fatalf("reverseBits: want: %d, got = %d", c.want, result)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	reverseBits(0543)
}