package solution1

func reverseWords(s string) string {
	if len(s) < 2 { return s }

	var stack = Stack{}
	stack.NewStack()
	ret := make([]rune, 0)

	for _, c := range s {
		if c != ' ' {
			stack.push(c)
		} else {
			if len(ret) != 0 {
				ret = append(ret, ' ')
			}
			for !stack.isEmpty() {
				ret = append(ret, stack.pop())
			}
		}
	}

	if !stack.isEmpty() {
		if len(ret) != 0 {
			ret = append(ret, ' ')
		}
		for !stack.isEmpty() {
			ret = append(ret, stack.pop())
		}
	}
	return string(ret)
}

type Stack struct {
	data []rune
}

func (s *Stack)NewStack() {
	s.data = make([]rune, 0)
}

func (s *Stack)push(c rune) {
	s.data = append(s.data, c)
}

func (s *Stack)pop() rune {
	pop := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data)-1]
	return pop
}

func (s *Stack)isEmpty() bool {
	return len(s.data) == 0
}