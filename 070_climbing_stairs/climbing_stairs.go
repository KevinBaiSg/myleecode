package _70_climbing_stairs

/*
	f(n) = f(n-1) + f(n-2)
	f(0) = f(1) = 1

	+ mem
*/

func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	x, y := 1, 2

	for i := 2; i < n; i++  {
		x, y = y, x + y
	}

	return y
}
