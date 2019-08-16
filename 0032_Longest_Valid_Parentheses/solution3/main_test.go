package solution3

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		in 		string
		want 	int
	}{
		{"", 0},
		{"(", 0},
		{")", 0},
		{"()", 2},
		{"()(())", 6},
		{"(()", 2},
	}

	for _, c := range cases {
		long := longestValidParentheses(c.in)
		if long != c.want {
			t.Fatalf("longestValidParentheses error: want: %d; got %d", c.want, long)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestValidParentheses("()(())")
	}
}
