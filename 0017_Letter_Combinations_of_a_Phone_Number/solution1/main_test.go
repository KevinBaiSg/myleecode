package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	letterCombinations("23")
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinations("23")
	}
}