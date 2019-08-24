package solution1

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 { return intervals}

	intsort := IntervalSort(intervals)
	sort.Sort(intsort)

	rets := make([][]int, 0)
	last := intsort[0]
	rets = append(rets, last)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= last[1] {
			last[1] = max(last[1], intervals[i][1])
		} else {
			last = intervals[i]
			rets = append(rets, last)
		}
	}

	return rets
}

func max(a, b int) int {
	if a < b { return b }
	return a
}

type IntervalSort [][]int

func (i IntervalSort) Len() int {
	return len(i)
}

func (i IntervalSort) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i IntervalSort) Less(a, b int) bool {
	return i[a][0] < i[b][0]
}

