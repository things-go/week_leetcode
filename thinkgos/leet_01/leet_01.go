package leet_01

type Item struct {
	Value    int
	Children []int
}

func Flat(s []Item, level int) []Item {
	out := make([]Item, 0, level*len(s))
	for _, v := range s {
		out = append(out, v)
	}

	return s
}
