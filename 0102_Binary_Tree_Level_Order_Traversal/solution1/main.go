package solution1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// 边界
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
	    return [][]int{
			{root.Val},
	    }
	}

	queue := make([]*TreeNode, 0)
	levels := make([][]int, 0)

	queue = append(queue, root)

	for len(queue) != 0 {
		n, level := len(queue), make([]int, 0)

		// 处理 level
		for i := 0; i < n; i++ {
			// 处理树
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left) // push()
			}

			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right) // push()
			}

			level = append(level, queue[0].Val)
			queue = queue[1:] // pop()
		}

		// 添加 level
		levels = append(levels, level)
	}

	return levels
}