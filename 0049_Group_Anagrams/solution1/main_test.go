package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		in 		[]string
		want 	[][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"ate","eat","tea"}, {"nat","tan"}, {"bat"}},},
	}

	for _, c := range cases {
		group := groupAnagrams(c.in)
		if len(group) != len(c.want) {
			t.Fatalf("groupAnagrams")
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	}
}