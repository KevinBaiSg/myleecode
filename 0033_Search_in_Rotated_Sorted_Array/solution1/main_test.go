package solution1

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct{
		in 		[]int
		target	int
		want	int
	}{
		{[]int{}, 1, -1},
		{[]int{3}, 3, 0},
		{[]int{1, 3}, 3, 1},
		{[]int{1, 3, 6, 8, 20, 34}, 6, 2},
		{[]int{1, 3, 6, 8, 20, 34}, 9, -1},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
	}
	for _, c := range cases {
		ret := search(c.in, c.target)
		if ret != c.want {
			t.Fatalf("Binary Search Failed; want: %d; got: %d", c.want, ret)
		}
	}
}

func TestRotateIndexSearch(t *testing.T) {
	cases := []struct{
		in 		[]int
		want	int
	}{
		{[]int{3}, 0},
		{[]int{1, 3}, 0},
		{[]int{20, 34, 1, 3, 6, 8}, 2},
		{[]int{6, 8, 20, 34, 1, 3}, 4},
		{[]int{1, 2, 3, 4, 5}, 0},
	}
	for i, c := range cases {
		ret := rotateIndexSearch(c.in, 0, len(c.in)-1)
		if ret != c.want {
			t.Fatalf("index: %d; Rotated Index Search Failed; want: %d; got: %d", i, c.want, ret)
		}
	}
}

func BenchmarkSerialIntArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search([]int{1, 3, 6, 8, 20, 34}, 3)
	}
}