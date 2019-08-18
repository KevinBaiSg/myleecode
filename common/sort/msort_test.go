package sort

import (
	"testing"

	. "github.com/KevinBaiSg/MyLeetCode/common"
)

func TestMergeSort(t *testing.T) {
	cases := []struct{
		in		[]int
		want 	[]int
	}{
		{[]int{}, []int{}},
		{[]int{-10, 200, 3, 6, 8, 30}, []int{-10, 3, 6, 8, 30, 200}},
	}
	for _, c := range cases {
		sort := MergeSort(c.in)
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

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort([]int{-10, 200, 3, 6, 8, 30})
	}
}