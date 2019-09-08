package solution1

import (
	"testing"

	"github.com/KevinBaiSg/MyLeetCode/common"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct{
		in		[]int
	}{
		{[]int{4,1,3,1}},
		{[]int{4,5,5,6}},
		{[]int{1,3,2,2,3,1}},
		{[]int{1,5,1,1,6,4}},
		{[]int{1,3,2,2,3,1}},
		{[]int{4,5,5,5,5,6,6,6}},
	}
	for _, c := range cases {
		wiggleSort(c.in)
		t.Logf("got: %s", common.SerialIntArray(c.in))
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++  {
		wiggleSort([]int{1,3,2,2,3,1})
	}
}