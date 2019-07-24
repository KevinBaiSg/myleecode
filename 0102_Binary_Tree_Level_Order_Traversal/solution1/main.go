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

	queue := Queue{}
	queue.NewQueue()
	levels := make([][]int, 0)

	queue.Enqueue(root)

	for !queue.isEmpty() {
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

type Queue struct {
	items []*TreeNode
}

func (q *Queue) NewQueue() *Queue {
	q.items = []*TreeNode{}
	return q
}

func (q *Queue) isEmpty() bool {
	if q.items == nil { return true }
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	if q.items == nil { return 0 }
	return len(q.items)
}

func (q *Queue) Peek() *TreeNode {
	if q.items == nil { return nil }

	if len(q.items) != 0 {
		return q.items[0]
	}

	return nil
}

func (q *Queue) Dequeue() *TreeNode {
	if q.items == nil { return nil }

	if len(q.items) == 0 { return nil }

	n := q.items[0]
	q.items = q.items[1:]
	return n
}

func (q *Queue) Enqueue(node *TreeNode) {
	if q.items == nil { return }

	q.items = append(q.items, node)
}