package reverse_integer

import "math"

/*
ref: https://leetcode.com/problems/reverse-integer/

Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
Input: 123
Output: 321

Example 2:
Input: -123
Output: -321

Example 3:
Input: 120
Output: 21

Note:
Assume we are dealing with an environment which could only store
integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 0
when the reversed integer overflows.
*/

/*
1.边界判断
2.特殊返回
	2^31-1=2147483647,-2^31=-2147483648
*/

func reverse(x int) int {
	if x > -9 && x < 10 {
		return x
	}

	ret := 0
	for x != 0 {
		m := x % 10
		if ret > math.MaxInt32/10 || (ret == math.MaxInt32/10 && m > 7) {
			return 0
		}
		if ret < math.MinInt32/10 || (ret == math.MinInt32/10 && m < -8) {
			return 0
		}

		ret = ret*10 + m
		x = x / 10
	}

	return ret
}
