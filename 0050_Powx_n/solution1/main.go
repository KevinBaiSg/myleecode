package solution1

func myPow(x float64, n int) float64 {
	if n == 0 { return 1 }
	if x == 1 || x == 0 { return x }

	const IntMax = int(^uint(0) >> 1)
	const IntMin = ^IntMax

	if n == IntMin {
		return myPow(x, n + 1) * myPow(x, - 1)
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	if n & 1 == 0 {
		return myPow(x * x, n / 2)
	} else {
		return x * myPow(x * x, n / 2)
	}
}