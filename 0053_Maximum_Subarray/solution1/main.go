package solution1

func maxSubArray(nums []int) int {
	if len(nums) == 1 { return nums[0] }

	sum := 0
	ans := nums[0]

	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}

		ans = max(sum, ans)
	}

	return ans
}

func max(a, b int) int {
	if a < b { return b }
	return a
}