package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		in 		[]int
		target	int
		want 	[][]int
	}{
		{[]int{1}, 2, [][]int{{1, 1}},},
		{[]int{2,3,6,7}, 7, [][]int{{2, 2, 3}, {7}},},
	}

	for _, c := range cases {
		combinationSums := combinationSum(c.in, c.target)
		if len(combinationSums) != len(c.want) {
			t.Fatalf("combinationSum")
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	//boilerplate()
}