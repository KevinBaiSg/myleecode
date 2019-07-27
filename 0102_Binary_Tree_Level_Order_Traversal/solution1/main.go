package solution1

import (
	"github.com/KevinBaiSg/myleecode/common/collections"
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
	if root.Left == nil && root.Right == nil {
		return [][]int{
			{root.Val},
		}
	}

	queue := collections.Queue{}
	queue.NewQueue()
	levels := make([][]int, 0)

	queue.Enqueue(root)

	for !queue.IsEmpty() {
		n, level := queue.Size(), make([]int, 0)

		// 处理 level
		for i := 0; i < n; i++ {
			// 处理树
			if queue.Peek().Left != nil {
				queue.Enqueue(queue.Peek().Left)
			}

			if queue.Peek().Right != nil {
				queue.Enqueue(queue.Peek().Right)
			}

			level = append(level, queue.Peek().Val)
			queue.Dequeue()
		}

		// 添加 level
		levels = append(levels, level)
	}

	return levels
}
