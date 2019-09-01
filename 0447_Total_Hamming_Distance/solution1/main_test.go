package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		in 		[]int
		want 	int
	}{
		{[]int{4,14,2}, 6},
	}

	for _, c := range cases {
		result := totalHammingDistance(c.in)
		if result != c.want {
			t.Fatalf("totalHammingDistance: want: %d, got = %d", c.want, result)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	totalHammingDistance([]int{4,14,2})
}