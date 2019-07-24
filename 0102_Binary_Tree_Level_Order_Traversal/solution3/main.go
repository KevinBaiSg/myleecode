package solution3

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
	if root == nil { return [][]int{} }

	// 使用 stack 处理 tree 访问顺序
	stack := Stack{}
	stack.NewStack()

	// 保存节点是否访问过的状态
	visited := make(map[*TreeNode]bool)

	// level 数组
	levels := make([][]int, 0)

	// 当前 stack pop 节点的 level
	level := 0 //

	if level >= stack.size() {
		levels = append(levels, make([]int, 0))
	}
	stack.push(root) // root 入栈

	levels[level] = append(levels[level], root.Val)

	for !stack.isEmpty() {
		last := stack.peek()
		// 入栈
		if last.Left != nil && visited[last.Left] != true {
			level ++
			if level >= len(levels) {
				levels = append(levels, make([]int, 0))
			}
			levels[level] = append(levels[level], last.Left.Val)
			stack.push(last.Left)
			visited[last.Left] = true

		} else if last.Right != nil  && visited[last.Right] != true {
			level ++
			if level >= len(levels) {
				levels = append(levels, make([]int, 0))
			}
			levels[level] = append(levels[level], last.Right.Val)
			stack.push(last.Right)
			visited[last.Right] = true
		} else { // 出栈
			stack.pop()
			level--
		}
	}

	return levels
}


type Stack struct {
	items []*TreeNode
}

func (s *Stack) NewStack() *Stack {
	s.items = []*TreeNode{}
	return s
}

func (s *Stack) pop() *TreeNode {
	if s.items == nil { return nil }

	if len(s.items) != 0 {
		item := s.items[len(s.items) - 1]
		s.items = s.items[:len(s.items) - 1]
		return item
	}
	return nil
}

func (s *Stack) peek() *TreeNode {
	if s.items == nil { return nil }

	if len(s.items) != 0 {
		item := s.items[len(s.items) - 1]
		return item
	}
	return nil
}

func (s *Stack) push(node *TreeNode) {
	if s.items == nil { return }

	s.items = append(s.items, node)
}

func (s *Stack) isEmpty() bool {
	if s.items == nil { return true }

	if len(s.items) == 0 {
		return true
	}

	return false
}

func (s *Stack) size() int {
	if s.items == nil { return 0 }

	return len(s.items)
}