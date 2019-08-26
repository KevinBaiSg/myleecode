package solution1

func hammingWeight(num uint32) int {
	temp := num - ((num >> 1) & 033333333333) - ((num >> 2) & 011111111111)
	return int(((temp + (temp >> 3)) & 030707070707) % 63)
}
