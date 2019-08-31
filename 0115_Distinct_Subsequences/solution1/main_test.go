package solution1

import (
	"testing"
)

func TestIsScramble(t *testing.T)  {
	cases := []struct {
		s 		string
		t 		string
		want 	int
	}{
		{"", "", 1},
		{"rabbbit", "rabbit", 3},
		{"babgbag", "bag", 5},
		{"adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc", "bcddceeeebecbc", 700531452},
	}
	for i, c := range cases {
		distinct := numDistinct(c.s, c.t)
		if distinct != c.want {
			t.Errorf("index: %d numDistinct failed; want: %d; got: %d", i, c.want, distinct)
		}
	}
}

func BenchmarkIsScramble(b *testing.B) {
	numDistinct("adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc", "bcddceeeebecbc")
}