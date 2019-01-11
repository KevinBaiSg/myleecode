package binarytreetilt

import (
	"fmt"
	"testing"
)

func TestPreorderTraversalRecursive(t *testing.T) {
	// trees := TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val:  2,
	// 		Left: nil,
	// 		Right: &TreeNode{
	// 			Val:   4,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 3,
	// 		Left: &TreeNode{
	// 			Val:   5,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: &TreeNode{
	// 			Val:   6,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 	},
	// }

	trees := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	ret := findTilt(&trees)
	fmt.Printf("%d\n", ret)

	return
}
