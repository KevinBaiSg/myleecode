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
	mem := make([]int, n)
	mem[0] = 1
	mem[1] = 2
	for i := 2; i < n; i++  {
		mem[i] = mem[i - 1] + mem[i - 2]
	}

	return mem[n - 1]
}
