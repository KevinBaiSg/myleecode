package solution2

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct{
		in 		[]int
		want	int
	} {
		{[]int{0,1,0,2,1,0,1,3,2,1,2,1}, 6},
	}

	for i, c := range cases {
		ret := trap(c.in)
		if ret != c.want {
			t.Fatalf("index:%d failed; want %d; got %d", i, c.want, ret)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trap([]int{0,1,0,2,1,0,1,3,2,1,2,1})
	}
}