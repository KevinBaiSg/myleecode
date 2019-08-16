package solution2

func longestValidParentheses(s string) int {
	if len(s) < 2 { return 0 }

	longest := 0
	stack := Stack{}
	stack.NewStack(len(s))
	stack.Push(-1)

	for i, c := range s {
		if c == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				longest = max(longest, i - stack.Seek())
			}
		}
	}

	return longest
}

func max(a, b int) int {
	if a < b { return b }
	return a
}

type Stack struct {
	items []int
}

func (s *Stack)NewStack(l int) {
	s.items = make([]int, 0, l)
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	item := s.Seek()
	if item == -1 {
		return -1
	}
	s.items = s.items[:len(s.items) - 1]
	return item
}

func (s *Stack) Seek() int {
	if len(s.items) == 0 { return -1 }
	item := s.items[len(s.items) - 1]
	return item
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
