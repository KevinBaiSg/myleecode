package tree

import (
	"strconv"
	"testing"

	. "github.com/KevinBaiSg/myleecode/common/makeTree"
)

func array2string(nodes []int) string {
	result := "["
	for _, node := range nodes {
		result += strconv.Itoa(node) + ","
	}
	return result + "]"
}

func TestPreOrderTraversalRecursion(t *testing.T) {
	cases := []struct{
		in 		string
		want 	string
	}{
		{"[3,9,20,null,null,15,7]", "[3,9,20,15,7,]"},
	}

	for _, c := range cases {
		var result = make([]int, 0)
		preOrderTraversalRecursion(MakeTree(c.in), &result)
		if array2string(result) !=  c.want {
			t.Fatalf("failed, want: %s; got: %s", c.want, array2string(result))
		}
		t.Logf("success result: %s", c.want)
	}
}

func TestInOrderTraversalRecursion(t *testing.T) {
	cases := []struct{
		in 		string
		want 	string
	}{
		{"[3,9,20,null,null,15,7]", "[9,3,15,20,7,]"},
	}

	for _, c := range cases {
		var result = make([]int, 0)
		inOrderTraversalRecursion(MakeTree(c.in), &result)
		if array2string(result) !=  c.want {
			t.Fatalf("failed, want: %s; got: %s", c.want, array2string(result))
		}
		t.Logf("success result: %s", c.want)
	}
}

func TestPostOrderTraversalRecursion(t *testing.T) {
	cases := []struct{
		in 		string
		want 	string
	}{
		{"[3,9,20,null,null,15,7]", "[9,15,7,20,3,]"},
	}

	for _, c := range cases {
		var result = make([]int, 0)
		postOrderTraversalRecursion(MakeTree(c.in), &result)
		if array2string(result) !=  c.want {
			t.Fatalf("failed, want: %s; got: %s", c.want, array2string(result))
		}
		t.Logf("success result: %s", c.want)
	}
}

func TestPreOrderTraversal(t *testing.T) {
	cases := []struct{
		in 		string
		want 	string
	}{
		{"[3,9,20,null,null,15,7]", "[3,9,20,15,7,]"},
	}

	for _, c := range cases {
		result := preOrderTraversal(MakeTree(c.in))
		if array2string(result) !=  c.want {
			t.Fatalf("failed, want: %s; got: %s", c.want, array2string(result))
		}
		t.Logf("success result: %s", c.want)
	}
}

func TestInOrderTraversal(t *testing.T) {
	cases := []struct{
		in 		string
		want 	string
	}{
		{"[3,9,20,null,null,15,7]", "[9,3,15,20,7,]"},
	}

	for _, c := range cases {
		result := inOrderTraversal(MakeTree(c.in))
		if array2string(result) !=  c.want {
			t.Fatalf("failed, want: %s; got: %s", c.want, array2string(result))
		}
		t.Logf("success result: %s", c.want)
	}
}

func TestPostOrderTraversal(t *testing.T) {
	cases := []struct{
		in 		string
		want 	string
	}{
		{"[3,9,20,null,null,15,7]", "[9,15,7,20,3,]"},
	}

	for _, c := range cases {
		result := postOrderTraversal(MakeTree(c.in))
		if array2string(result) !=  c.want {
			t.Fatalf("failed, want: %s; got: %s", c.want, array2string(result))
		}
		t.Logf("success result: %s", c.want)
	}
}

func TestLayerTraver(t *testing.T) {
	cases := []struct{
		in 		string
		want 	string
	}{
		{"[3,9,20,null,null,15,7]", "[3,9,20,15,7,]"},
	}

	for _, c := range cases {
		result := layerTraver(MakeTree(c.in))
		if array2string(result) !=  c.want {
			t.Fatalf("failed, want: %s; got: %s", c.want, array2string(result))
		}
		t.Logf("success result: %s", c.want)
	}
}
