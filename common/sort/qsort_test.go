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
		{[]int{}, []int{}},
		{[]int{-10, 200, 3, 6, 8, 30}, []int{-10, 3, 6, 8, 30, 200}},
		{[]int{-1,0,1,2,-1,-4}, []int{-4,-1,-1,0,1,2}},
	}
	for _, c := range cases {
		QuickSort(c.in)
		for i := 0; i < len(c.in); i++ {
			if c.in[i] != c.want[i] {
				t.Fatalf("sort error, want: %s; got: %s", SerialIntArray(c.want), SerialIntArray(c.in))
			}
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort([]int{-10, 200, 3, 6, 8, 30})
	}
}
