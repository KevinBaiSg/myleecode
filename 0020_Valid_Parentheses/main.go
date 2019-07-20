package Valid_Parentheses

/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/

func isValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	runes := []rune(s)
	stack := make([]rune, len(runes))
	len := 0
	for _, characters := range runes {
		if pair, ok := m[characters]; ok && len > 0 && stack[len-1] == pair { // pop
			len--
		} else { // push
			stack[len] = characters
			len++
		}
	}

	return len == 0
}
