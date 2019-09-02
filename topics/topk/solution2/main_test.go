package solution2

import (
	"sort"
	"testing"

	. "github.com/KevinBaiSg/MyLeetCode/common"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct{
		in		[]int
		k 		int
		want 	[]int
	}{
		{[]int{5, 3, 7, 2, 4, 9, 19, 14, 73, 88, 6, 64, 17, 52}, 4, []int{52, 64, 73, 88}},
	}
	for _, c := range cases {
		tops := topK(c.in, c.k)
		if len(c.want) != len(tops) {
			t.Fatalf("topK error, want: %s; got: %s", SerialIntArray(c.want), SerialIntArray(tops))
		}

		sort.Ints(tops)
		for i := 0; i < len(tops); i++ {
			if tops[i] != c.want[i] {
				t.Fatalf("sort error, want: %s; got: %s", SerialIntArray(c.want), SerialIntArray(tops))
			}
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		topK([]int{5, 3, 7, 2, 4, 9, 19, 14, 73, 88, 6, 64, 17, 52}, 4)
	}
}
