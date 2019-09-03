package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		J 		string
		S 		string
		want	int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for _, c := range cases {
		num := numJewelsInStones(c.J, c.S)
		if num != c.want {
			t.Fatalf("numJewelsInStones: want: %d, got = %d", c.want, num)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numJewelsInStones("aA", "aAAbbbb")
	}
}