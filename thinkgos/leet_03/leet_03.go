package leet_03

type Dept struct {
	Id   int
	Pid  int
	Name string
}

type DeptLabel struct {
	*Dept
	Children []DeptLabel
}

func toDeptMap(rows []Dept) map[int][]*Dept {
	mp := make(map[int][]*Dept)
	for i := 0; i < len(rows); i++ {
		v := &rows[i]
		mp[v.Pid] = append(mp[v.Pid], v)
	}
	return mp
}

func toTree(rows []Dept, pid int) []DeptLabel {
	mp := toDeptMap(rows)
	rt, ok := mp[pid]
	if !ok || len(rt) == 0 {
		return []DeptLabel{}
	}
	root := make([]DeptLabel, 0, len(rt))
	for _, itm := range rt {
		root = append(root, deepChildrenDept(mp, DeptLabel{Dept: itm}))
	}
	return root
}

func deepChildrenDept(mp map[int][]*Dept, itm DeptLabel) DeptLabel {
	children, ok := mp[itm.Id]
	if ok {
		itm.Children = make([]DeptLabel, 0, len(children))
		for _, v := range children {
			itm.Children = append(itm.Children, deepChildrenDept(mp, DeptLabel{Dept: v}))
		}
	} else {
		itm.Children = []DeptLabel{}
	}
	return itm
}
