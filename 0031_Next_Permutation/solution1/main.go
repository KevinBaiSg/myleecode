package solution1

func nextPermutation(nums []int)  {
	if len(nums) < 2 { return }

	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[j], nums[i] = nums[i], nums[j]
	}

	reverse(nums, i + 1)
}

func reverse(nums []int, start int) {
	if len(nums) < 2 { return }

	i, j := start, len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++; j--
	}
}