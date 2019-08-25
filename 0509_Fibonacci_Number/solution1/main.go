package solution1

func fib(N int) int {
	if N < 2 { return N }
	return fib(N-1) + fib(N-2)
}
