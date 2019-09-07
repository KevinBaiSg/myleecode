package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		in 		string
		want 	bool
	}{
		{"0P", false},
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for i, c := range cases {
		ret := isPalindrome(c.in)
		if ret != c.want {
			t.Fatalf("index: %d; isPalindrome: want: %t, got = %t", i, c.want, ret)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("A man, a plan, a canal: Panama")
	}
}