package solution1

import (
	"testing"

	. "github.com/KevinBaiSg/myleecode/common"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct{
		in		[]int
		target 	int
		want	[]int
	}{
		{[]int{5,7,7,8,8,10}, 8, []int{3,4}},
	}
	for i, c := range cases {
		ret := searchRange(c.in, c.target)
		if ret[0] != c.want[0] || ret[1] != c.want[1] {
			t.Fatalf("case %d Failed; want: %s; got: %s", i, SerialIntArray(c.want), SerialIntArray(ret))
		}
	}

}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchRange([]int{5,7,7,8,8,10}, 8)
	}
}