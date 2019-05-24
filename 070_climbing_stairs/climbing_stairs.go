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
	mem := make([]int, 3)
	mem[0] = 1
	mem[1] = 2
	/*
			  mem[0]

		mem[1]		mem[2]
	*/

	for i := 2; i < n; i++  {
		mem[i % 3] = mem[(i - 1) % 3] + mem[(i - 2) % 3]
	}

	return mem[(n - 1) % 3]
}
