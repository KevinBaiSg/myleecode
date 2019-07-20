package majorityelement

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maj, count := nums[0], 0
	for _, num := range nums {
		if count == 0 {
			maj, count = num, 1
		} else if maj == num {
			count++
		} else {
			count--
		}
	}
	return maj
}
