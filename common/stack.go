package common

type Stack struct {
	items []*TreeNode
}

func (s *Stack) NewStack() *Stack {
	s.items = []*TreeNode{}
	return s
}

func (s *Stack) Pop() *TreeNode {
	if s.items == nil { return nil }

	if len(s.items) != 0 {
		item := s.items[len(s.items) - 1]
		s.items = s.items[:len(s.items) - 1]
		return item
	}
	return nil
}

func (s *Stack) Peek() *TreeNode {
	if s.items == nil { return nil }

	if len(s.items) != 0 {
		item := s.items[len(s.items) - 1]
		return item
	}
	return nil
}

func (s *Stack) Push(node *TreeNode) {
	if s.items == nil { return }

	s.items = append(s.items, node)
}

func (s *Stack) IsEmpty() bool {
	if s.items == nil { return true }

	if len(s.items) == 0 {
		return true
	}

	return false
}

func (s *Stack) Size() int {
	if s.items == nil { return 0 }

	return len(s.items)
}