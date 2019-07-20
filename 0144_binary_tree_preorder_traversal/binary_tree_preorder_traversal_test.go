package binarytreepreordertraversal

import (
	"fmt"
	"testing"
)

func TestPreorderTraversalRecursive(t *testing.T) {
	trees := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}

	rets := preorderTraversalRecursive(&trees)
	for _, ret := range rets {
		fmt.Printf("%2d", ret)
	}
	fmt.Println()

	rets = preorderTraversalIteration(&trees)
	for _, ret := range rets {
		fmt.Printf("%2d", ret)
	}
	fmt.Println()
	return
}
