package common

import "testing"

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
	}
	for i, c := range cases {
		ret := BinarySearch(c.in, c.target)
		if ret != c.want {
			t.Fatalf("index=%d Binary Search Failed; want: %d; got: %d", i, c.want, ret)
		}
	}
}

func BenchmarkSerialIntArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch([]int{1, 3, 6, 8, 20, 34}, 3)
	}
}