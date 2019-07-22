package _071_Simplify_Path

import "strings"

func simplifyPath(path string) string {
	var stack = make([]string , 0, 0)
	var paths = strings.Split(path, "/")

	for _, p := range paths {
		if p == "." || p == "" { continue } // 不处理
		if p == ".." { // 出栈
			if len(stack) > 0 {
				stack = stack[:len(stack) - 1]
			}
			continue
		}

		stack = append(stack, p) // 入栈
	}

	if len(stack) == 0 { return "/" }

	s := ""
	for _, sk := range stack {
		s += "/" + sk
	}

	return s
}