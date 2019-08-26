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
		want 	int
	}{
		{00, 0},
		{01, 1},
		{045, 3},
		{076355, 11},
		{040201, 3},
	}

	for index, c := range cases {
		got := hammingWeight(c.in)
		if got != c.want {
			t.Fatalf("index: %d; hammingWeight: want: %d, got = %d", index, c.want, got)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hammingWeight(45)
	}
}