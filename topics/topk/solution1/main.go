package solution1

import "sort"

func topK(nums []int, k int) []int {
	if len(nums) <= k { return nums }
	sort.Ints(nums) // O(n log n)
	return nums[len(nums)-k:]
}
