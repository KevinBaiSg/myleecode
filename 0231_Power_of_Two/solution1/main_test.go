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
		in 		int
		want 	bool
	}{
		{00, false},
		{01, true},
		{03, false},
		{04000, true},
		{06000, false},
	}

	for index, c := range cases {
		got := isPowerOfTwo(c.in)
		if got != c.want {
			t.Fatalf("index: %d; hammingWeight: want: %t, got = %t", index, c.want, got)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	isPowerOfTwo(04000)
}