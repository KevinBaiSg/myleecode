package binarytreetilt

import "math"

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

var tilt int

func tiltTraversal(node *TreeNode) int {
	left, right := 0, 0
	if node == nil {
		return 0
	}
	if node.Left != nil {
		left = tiltTraversal(node.Left)
	}

	if node.Right != nil {
		right = tiltTraversal(node.Right)
	}

	tilt = tilt + int(math.Abs(float64(left-right)))

	return left + right + node.Val
}

func findTilt(root *TreeNode) int {
	tiltTraversal(root)
	return tilt
}
