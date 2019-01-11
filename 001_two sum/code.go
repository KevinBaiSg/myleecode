package two_sum

/*
	思路：暴力破解，两个 for 嵌套
*/
func twoSumS1(nums []int, target int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if num1+num2 == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/*
	1.排序
	2.排序情况下处理和的问题：比如使用前后两个指针向中间靠拢
*/
// func twoSumS2(nums []int, target int) []int {
// 	sort.Ints(nums)
// 	i := 0
// 	j := len(nums) - 1
// 	for i < j {
// 		if nums[i]+nums[j] == target {
// 			return []int{i, j}
// 		} else if nums[i]+nums[j] > target {
// 			j--
// 		} else {
// 			i++
// 		}
// 	}

// 	return []int{}
// }

/*
	1. A + B = T
	2. for 循环中从map中查找B
	3. for 循环中把数组记录到map中 map[B]=i
	问题：map[B] B 可能会有重复变量
*/
func twoSumS3(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		j, ok := m[target-num]
		if ok && i != j {
			return []int{i, j}
		}
		m[num] = i
	}
	return []int{}
}
