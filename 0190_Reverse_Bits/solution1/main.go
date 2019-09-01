package solution1

func reverseBits(num uint32) uint32 {
	if num == 0 { return num }

	var reverse uint32 = 0

	for i := 0; i < 32; i++ {
		reverse = reverse << 1 | num & 1
		num = num >> 1
	}
	return reverse
}