package solution1

import (
	"testing"

	. "github.com/KevinBaiSg/myleecode/common"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		in		[]int
		want 	[]int
	}{
		{[]int{2, 3, 4, 1, 6}, []int{2, 3, 4, 6, 1}},
		{[]int{1, 2, 3, 4, 6}, []int{1, 2, 3, 6, 4}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, c := range cases {
		nextPermutation(c.in)
		if len(c.want) != len(c.in) {
			t.Fatalf("sort error, want: %s; got: %s", SerialIntArray(c.want), SerialIntArray(c.in))
		}
		for i := 0; i < len(c.in); i++ {
			if c.in[i] != c.want[i] {
				t.Fatalf("sort error, want: %s; got: %s", SerialIntArray(c.want), SerialIntArray(c.in))
			}
		}
	}

}

func BenchmarkIsScramble(b *testing.B) {
	nextPermutation([]int{1, 2, 3, 6, 4})
}