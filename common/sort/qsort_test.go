package sort

import (
	"testing"

	. "github.com/KevinBaiSg/myleecode/common"
)

func TestQuickSort(t *testing.T) {
	cases := []struct{
		in		[]int
		want 	[]int
	}{
		{[]int{1, 2, 4}, []int{2, 4, 3}},
	}
	for _, c := range cases {
		sort := QuickSort(c.in)
		if len(c.want) != len(sort) {
			t.Fatalf("sort error, want: %s; got: %s", SerialIntArray(c.want), SerialIntArray(sort))
		}
		for i := 0; i < len(sort); i++ {
			if sort[i] != c.want[i] {
				t.Fatalf("sort error, want: %s; got: %s", SerialIntArray(c.want), SerialIntArray(sort))
			}
		}
	}
}
