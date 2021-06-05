package leet_06

func SingleNumber(ss []int) int {
	v := 0
	for _, s := range ss {
		v ^= s
	}
	return v
}
