package solution2

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct{
		in 		int
		want 	int
	} {
		{0, 0},
		{1, 1},
		{2, 1},
		{18, 2584},
	}

	for _, c := range cases {
		got := fib(c.in)
		if got != c.want {
			t.Fatalf("fib error; want: %d; got: %d", c.want, got)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(18)
	}
}