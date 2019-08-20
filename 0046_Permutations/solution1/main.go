package solution1

func permute(nums []int) [][]int {
	if len(nums) == 0 { return [][]int{} }
	if len(nums) == 1 { return [][]int{ []int{ nums[0] } }}

	result := make([][]int, 0)

	result = backtrack(nums, 0, result)

	return result
}

func backtrack(nums []int, left int, result [][]int) (ret [][]int) {
	ret = result
	if left > len(nums) - 2 {
		ret = append(ret, nums)
		return
	}

	for i := left; i < len(nums); i++ {
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		newNums[left], newNums[i] = newNums[i], newNums[left]
		ret = backtrack(newNums, left + 1, ret)
	}

	return ret
}