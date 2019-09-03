package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct{
		in		[]int
		want 	string
	}{
		{[]int{0, 0}, "0"},
		{[]int{10, 2}, "210"},
		{[]int{3,30,34,5,9}, "9534330"},
	}
	for _, c := range cases {
		largest := largestNumber(c.in)
		if len(c.want) != len(largest) {
			t.Fatalf("largestNumber error, want: %s; got: %s", c.want, largest)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestNumber([]int{3,30,34,5,9})
	}
}