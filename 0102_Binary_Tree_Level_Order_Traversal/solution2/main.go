package solution2

import (
	"github.com/KevinBaiSg/myleecode/common/makeTree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *makeTree.TreeNode) [][]int {
	// 边界
	if root == nil {
		return [][]int{}
	}

	levels := make([][]int, 0)

	dfs(root, 0, &levels)

	return levels
}

func dfs(root *makeTree.TreeNode, level int, levels *[][]int) {
	if root == nil { return }

	if level >= len(*levels) {
		*levels = append(*levels, []int{root.Val})
	} else {
		(*levels)[level] = append((*levels)[level], root.Val)
	}

	dfs(root.Left, level + 1, levels)
	dfs(root.Right, level + 1, levels)
}