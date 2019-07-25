package solution1

import (
	. "github.com/KevinBaiSg/myleecode/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	if root == nil { return nil }
	next(root, 0)
	return root
}

func next(root *TreeNode, value int) int {
	if root != nil {
		value = next(root.Right, value)
		value = root.Val + value
		root.Val = value
		value = next(root.Left, value)
	}
	return value
}
