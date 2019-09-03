package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		in 		string
		want 	string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	}

	for _, c := range cases {
		reverse := reverseWords(c.in)
		if reverse != c.want {
			t.Fatalf("reverseWords: in = %s, want: %s, got = %s", c.in, c.want, reverse)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWords("Let's take LeetCode contest")
	}
}