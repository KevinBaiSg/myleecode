package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		in 		[]int
		want 	int
	}{
		{[]int{-2,1,-3,4,-1,2,1,-5,4}, 6},
	}

	for _, c := range cases {
		max := maxSubArray(c.in)
		if max != c.want {
			t.Fatalf("maxSubArray")
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})
}