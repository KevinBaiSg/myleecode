package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	//boilerplate()
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(3)
	}
}