package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct{
		nums	[]int
		k 		int
		want	int
	}{
		//{[]int{-1,2,0}, 1, 2},
		{[]int{-1,2,0}, 3, -1},
		//{[]int{3,2,1,5,6,4}, 2, 5},
		//{[]int{3,2,3,1,2,4,5,5,6}, 4, 4},
	}
	for i, c := range cases {
		top := findKthLargest(c.nums, c.k)
		if top != c.want {
			t.Fatalf("index: %d findKthLargest error: want: %d; got: %d", i, c.want, top)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	//findKthLargest()
}