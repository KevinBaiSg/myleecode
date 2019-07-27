package collections

import (
	. "github.com/KevinBaiSg/myleecode/common"
)

type Queue struct {
	items []*TreeNode
}

func (q *Queue) NewQueue() *Queue {
	q.items = []*TreeNode{}
	return q
}

func (q *Queue) IsEmpty() bool {
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