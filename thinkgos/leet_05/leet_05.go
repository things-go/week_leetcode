package leet_05

import (
	"sort"
)

func UniqueSortList(ss []int) []int {
	sort.Ints(ss)
	idx := 0
	for i := 1; i < len(ss); i++ {
		if ss[idx] != ss[i] {
			idx++
			ss[idx] = ss[i]
		}
	}
	return ss[:idx+1]
}
