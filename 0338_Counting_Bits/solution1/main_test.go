package solution1

import (
	"testing"

	. "github.com/KevinBaiSg/MyLeetCode/common"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct{
		in		int
		want 	[]int
	}{
		{0, []int{0}},
		{2, []int{0,1,1}},
		{5, []int{0,1,1,2,1,2}},
	}
	for _, c := range cases {
		counts := countBits(c.in)
		for i := 0; i < len(counts); i++ {
			if counts[i] != c.want[i] {
				t.Fatalf("sort error, want: %s; got: %s", SerialIntArray(c.want), SerialIntArray(counts))
			}
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countBits(5)
	}
}