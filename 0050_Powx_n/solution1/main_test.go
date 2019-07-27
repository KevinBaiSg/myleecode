package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	myPow(10, 10)
}

func BenchmarkIsScramble(b *testing.B) {
	myPow(10, 100)
}