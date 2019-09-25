package solution1

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	cases := []struct {
		s    string
		p    string
		want bool
	}{
		{"aa", 		"a", 		false},
		{"aa", 		"*", 		true},
		{"cb", 		"?a", 		false},
		{"acdcb", 	"a*c?b", 	false},
		{"adceb", 	"*a*b", 		true},
	}

	for index, c := range cases {
		ret := isMatch(c.s, c.p)
		if ret != c.want {
			t.Fatalf("index= %d isMatch error; want: %t; got: %t ", index, c.want, ret)
		}
	}
}

func BenchmarkIsMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isMatch("adceb", "*a*b")
	}
}
