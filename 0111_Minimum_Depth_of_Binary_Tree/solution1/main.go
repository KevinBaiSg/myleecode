package solution1

import (
	"github.com/KevinBaiSg/MyLeetCode/common/makeTree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *makeTree.TreeNode) int {
	//     边界
	if root == nil { return 0 }
	if root.Left == nil && root.Right == nil { return 1 }

	// nil 会影响 depth 计算
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}

	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	return 1 + min(minDepth(root.Left), minDepth(root.Right))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
