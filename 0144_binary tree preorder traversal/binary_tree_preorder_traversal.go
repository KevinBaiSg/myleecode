package binarytreepreordertraversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode ..
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversalRecursive(root *TreeNode) []int {
	rets := []int{}
	if root == nil {
		return rets
	}

	rets = append(rets, root.Val)

	if root.Left != nil {
		ret := preorderTraversalRecursive(root.Left)
		rets = append(rets, ret...)
	}

	if root.Right != nil {
		ret := preorderTraversalRecursive(root.Right)
		rets = append(rets, ret...)
	}

	return rets
}

func preorderTraversalIteration(root *TreeNode) []int {
	rets := []int{}
	if root == nil {
		return rets
	}

	stack := []*TreeNode{}
	p := root
	for len(stack) != 0 || p != nil {
		for p != nil {
			rets = append(rets, p.Val)
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	return rets
}
