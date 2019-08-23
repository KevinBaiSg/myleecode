package sort

import (
	"testing"
)

func TestInterfaceSort(t *testing.T) {
	cases := []struct{
		in		string
		want 	string
	}{
		{"qewrtysgfda", "adefgqrstwy"},
	}
	for _, c := range cases {
		sort := InterfaceSort(c.in)
		if len(c.want) != len(sort) {
			t.Fatalf("sort error, want: %s; got: %s", c.want, sort)
		}
	}
}

func BenchmarkInterfaceSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HeapSort([]int{-10, 200, 3, 6, 8, 30})
	}
}