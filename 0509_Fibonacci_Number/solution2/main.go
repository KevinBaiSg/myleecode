package solution2

func fib(N int) int {
	if N < 2 { return N }
	N1, N2 := 0, 1
	for i := 2; i < N; i++ {
		N2, N1 = N1+N2, N2
	}
	return N1+N2
}
