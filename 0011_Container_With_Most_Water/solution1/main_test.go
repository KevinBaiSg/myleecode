package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		in 		[]int
		want 	int
	}{
		{[]int{1,8,6,2,5,4,8,3,7}, 49},
	}

	for _, c := range cases {
		result := maxArea(c.in)
		if result != c.want {
			t.Fatalf("maxArea: want: %d, got = %d", c.want, result)
		}
	}

}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		_ = maxArea([]int{1,8,6,2,5,4,8,3,7})
	}
}