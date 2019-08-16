package solution3

func longestValidParentheses(s string) int {
	if len(s) < 2 { return 0 }
	left, right := 0, 0
	longest := 0

	for _, c := range s {
		if c == '(' {
			left++
		} else {
			right++
		}
		if right == left {
			longest = max(longest, 2 * right)
		} else if right >= left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if right == left {
			longest = max(longest, 2 * right)
		} else if left >= right {
			left, right = 0, 0
		}
	}

	return longest
}

func max(a, b int) int {
	if a < b { return b }
	return a
}
