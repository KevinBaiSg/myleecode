package solution3

import (
	"github.com/KevinBaiSg/MyLeetCode/common/collections"
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

func levelOrder(root *makeTree.TreeNode) [][]int {
	// 边界
	if root == nil { return [][]int{} }

	// 使用 stack 处理 tree 访问顺序
	stack := collections.Stack{}
	stack.NewStack()

	// 保存节点是否访问过的状态
	visited := make(map[*makeTree.TreeNode]bool)

	// level 数组
	levels := make([][]int, 0)

	// 当前 stack pop 节点的 level
	level := 0 //

	if level >= stack.Size() {
		levels = append(levels, make([]int, 0))
	}
	stack.Push(root) // root 入栈

	levels[level] = append(levels[level], root.Val)

	for !stack.IsEmpty() {
		last := stack.Peek()
		// 入栈
		if last.Left != nil && visited[last.Left] != true {
			level ++
			if level >= len(levels) {
				levels = append(levels, make([]int, 0))
			}
			levels[level] = append(levels[level], last.Left.Val)
			stack.Push(last.Left)
			visited[last.Left] = true

		} else if last.Right != nil  && visited[last.Right] != true {
			level ++
			if level >= len(levels) {
				levels = append(levels, make([]int, 0))
			}
			levels[level] = append(levels[level], last.Right.Val)
			stack.Push(last.Right)
			visited[last.Right] = true
		} else { // 出栈
			stack.Pop()
			level--
		}
	}

	return levels
}
