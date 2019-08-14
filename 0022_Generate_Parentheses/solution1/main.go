package solution1

var parentheses []string

func generateParenthesis(n int) []string {
	if n == 0 { return []string{} }
	if n == 1 { return []string{"()"} }

	parentheses = make([]string, 0)
	backtrack("", 0, 0, n)
	return parentheses
}

func backtrack(cur string, left int, right int, n int) {
	if left > right {
		if left < n {
			backtrack(cur + "(", left + 1, right, n)
		}
		backtrack(cur + ")", left, right + 1, n)
	} else { // ==
		if left == n {
			parentheses = append(parentheses, cur)
		} else {
			backtrack(cur + "(", left + 1, right, n)
		}
	}
}