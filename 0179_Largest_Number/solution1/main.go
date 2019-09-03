package solution1

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	sort.Sort(Largest(strs))

	ret := ""
	for _, str := range strs {
		if ret == "0" && str == "0" { continue } // 特殊情况 "00..."
		ret += str
	}

	return ret
}

type Largest []string

func (l Largest)Len() int {
	return len(l)
}

func (l Largest)Less(i, j int) bool {
	if l[i] + l[j] > l[j] + l[i] {
		return true
	}
	return false
}

func (l Largest)Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

