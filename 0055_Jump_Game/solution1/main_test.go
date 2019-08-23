package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		in 		[]int
		want 	bool
	}{
		{[]int{2,3,1,1,4}, true},
	}

	for _, c := range cases {
		max := canJump(c.in)
		if max != c.want {
			t.Fatalf("canJump failed")
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canJump([]int{2,3,1,1,4})
	}
}