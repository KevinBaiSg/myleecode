package solution1

func totalHammingDistance(nums []int) int {
	if len(nums) < 2 { return 0 }

	total := 0
	mask := 1
	for i := 0; i < 32; i++ {
		mask0 := 0

		for _, num := range nums {
			if num & mask == 0 {
				mask0++
			}
		}

		total += mask0 * (len(nums) - mask0)

		mask <<= 1
	}

	return total
}
