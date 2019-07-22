package _071_Simplify_Path

import "testing"

func TestSimplifyPath(t *testing.T) {
	cases := []struct {
		in 		string
		want 	string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
	}

	for _, c := range cases {
		path := simplifyPath(c.in)
		if path != c.want {
			t.Fatalf("simplifyPath: in = %s, want: %s, got = %s", c.in, c.want, path)
		}
	}
}

func BenchmarkSimplifyPath( b *testing.B)  {
	for i := 0; i < b.N ; i++ {
		simplifyPath("/a/./b/../../c/")
	}
}