package solution1

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 { return [][]int{} }
	qsort(candidates)
	//sort.Ints(candidates)
	return dfs(candidates, 0, target, []int{}, [][]int{})
}

func dfs(candidates []int, index int, target int, path []int, result [][]int) (ret [][]int) {
	ret = result
	if index > len(candidates) - 1 { return }
	if target < 0 { return }
	if target == 0 {
		ret = append(ret, path)
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i - 1] { continue } // 去重
		newPath := make([]int, len(path))
		copy(newPath, path)
		newPath = append(newPath, candidates[i])
		ret = dfs(candidates, i, target - candidates[i], newPath, ret)
	}
	return
}

func qsort(candidates []int) {
	if len(candidates) < 2 { return }

	pivot := candidates[0]
	l, r := 0, len(candidates) - 1
	for i := 1; i <= r; {
		if candidates[i] < pivot {
			candidates[l], candidates[i] = candidates[i], candidates[l]
			l++; i++
		} else {
			candidates[r], candidates[i] = candidates[i], candidates[r]
			r--
		}
	}

	qsort(candidates[:l])
	qsort(candidates[l+1:])
}
