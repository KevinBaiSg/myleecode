package solution1

import (
	. "github.com/KevinBaiSg/MyLeetCode/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 { return nil }
	if len(nums) == 1 { return &TreeNode{ Val: nums[0], Left: nil, Right: nil }}

	return bfs(nums, 0, len(nums)-1)
}

func bfs(nums []int, l, r int) *TreeNode {
	if l == r { return  &TreeNode{ Val: nums[l], Left: nil, Right: nil } }

	m := l + (r - l) / 2

	root := TreeNode{ Val: nums[m], Left: nil, Right: nil }
	if m-1 >= l {
		root.Left = bfs(nums, l, m-1)
	}
	if m+1 <= r {
		root.Right = bfs(nums, m+1, r)
	}

	return &root
}