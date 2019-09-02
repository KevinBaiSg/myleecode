package solution2

func topK(nums []int, k int) []int {
	if len(nums) <= k { return nums }

	for i := 0; i < k; i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums[len(nums)-k:]
}
