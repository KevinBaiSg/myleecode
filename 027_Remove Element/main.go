package _27_Remove_Element

func removeElement(nums []int, val int) int {
	var (
		removed = 0
		length = 0
	)
	for index, num := range nums {
		if num == val {
			removed ++
		} else {
			length ++
			if removed > 0 {
				nums[index - removed] = nums[index]
			}
		}
	}
	return length
}
