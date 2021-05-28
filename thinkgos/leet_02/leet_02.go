package leet_02

import (
	"sort"
)

type Int map[int]struct{}

func New(items ...int) Int {
	ss := Int{}
	for _, v := range items {
		ss[v] = struct{}{}
	}
	return ss
}

// Difference returns a set of objects that are not in s2
// For example:
// s = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s.Difference(s2) = {a3}
// s2.Difference(s) = {a4, a5}
func (s Int) Difference(s2 Int) Int {
	result := New()
	for key := range s {
		_, contained := s2[key]
		if !contained {
			result[key] = struct{}{}
		}
	}
	return result
}

// List returns the contents as a sorted int8 slice.
func (s Int) List() []int {
	res := make([]int, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	sort.Sort(sort.IntSlice(res))
	return res
}
