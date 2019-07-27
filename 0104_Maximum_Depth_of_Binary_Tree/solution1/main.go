package solution1

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
func maxDepth(root *makeTree.TreeNode) int {
	// 边界
	if root == nil { return 0 }
	if root.Left == nil && root.Right == nil { return 1 }

	// queue
	var (
		max = 0
		q = Queue{}
	)

	q.NewQueue()

	q.enqueue(root)
	for !q.isEmpty() {
		max++
		n := q.size()
		for i := 0; i < n; i++ {
			node := q.dequeue()
			if node.Left != nil {
				q.enqueue(node.Left)
			}
			if node.Right != nil {
				q.enqueue(node.Right)
			}
		}
	}

	return max
}

type Queue struct {
	items []*makeTree.TreeNode
}

func (q *Queue) NewQueue() *Queue {
	q.items = make([]*makeTree.TreeNode, 0)
	return q
}

// enqueue
func (q *Queue) enqueue(node *makeTree.TreeNode) {
	if q.items == nil { return }

	q.items = append(q.items, node)
}

// dequeue
func (q *Queue) dequeue() *makeTree.TreeNode {
	if q.items == nil { return nil }
	if len(q.items) == 0 { return nil }

	n := q.items[0]
	q.items = q.items[1:]
	return n
}

// size
func (q *Queue) size() int {
	if q.items == nil { return 0 }
	return len(q.items)
}

// isEmpty
func (q *Queue) isEmpty() bool {
	if q.items == nil { return true }
	return len(q.items) == 0
}

// peek
func (q *Queue) peek() *makeTree.TreeNode {
	if q.items == nil { return nil }
	if len(q.items) == 0 { return nil }
	return q.items[0]
}
