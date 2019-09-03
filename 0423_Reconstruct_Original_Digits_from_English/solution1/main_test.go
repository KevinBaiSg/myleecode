package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		in 		string
		want 	string
	}{
		{"owoztneoer", "012"},
		{"fviefuro", "45"},
	}

	for _, c := range cases {
		digits := originalDigits(c.in)
		if digits != c.want {
			t.Fatalf("originalDigits: in = %s, want: %s, got = %s", c.in, c.want, digits)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		originalDigits("owoztneoer")
	}
}