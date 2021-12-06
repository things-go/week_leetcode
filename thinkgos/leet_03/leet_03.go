package leet_03

type Dept struct {
	Id   int
	Pid  int
	Name string
}

type DeptTree struct {
	*Dept
	Children []DeptTree
}

func toDeptMap(rows []Dept) map[int][]*Dept {
	mp := make(map[int][]*Dept)
	for i := 0; i < len(rows); i++ {
		v := &rows[i]
		mp[v.Pid] = append(mp[v.Pid], v)
	}
	return mp
}

func toTree(rows []Dept, pid int) []DeptTree {
	mp := toDeptMap(rows)
	rt, ok := mp[pid]
	if !ok || len(rt) == 0 {
		return []DeptTree{}
	}
	root := make([]DeptTree, 0, len(rt))
	for _, itm := range rt {
		root = append(root, deepChildrenDept(mp, DeptTree{Dept: itm}))
	}
	return root
}

func deepChildrenDept(mp map[int][]*Dept, itm DeptTree) DeptTree {
	children, ok := mp[itm.Id]
	if ok {
		itm.Children = make([]DeptTree, 0, len(children))
		for _, v := range children {
			itm.Children = append(itm.Children, deepChildrenDept(mp, DeptTree{Dept: v}))
		}
	} else {
		itm.Children = []DeptTree{}
	}
	return itm
}

/**
function toTree(arr, pid) {
    let tree = [];
    let hash = {};
    arr.forEach((item) => {
        const id = item.id;
        item["children"] = hash[id] = [];
        // 判断是否为父节点
        if (item.pid === pid) {
            tree.push(item);
        } else {
            const parentId = item.pid;
            if (hash[parentId]) {
                hash[parentId].push(item);
            } else {
                hash[parentId] = [];
                hash[parentId].push(item);
            }
        }
    });
    console.log(hash)
    return tree;
}
*/

type DeptTree2 struct {
	*Dept
	Children *[]*DeptTree2
}

func toTree2(rows []Dept, pid int) []*DeptTree2 {
	mp := make(map[int]*[]*DeptTree2) // 孩子
	root := make([]*DeptTree2, 0, 10)
	for i := 0; i < len(rows); i++ {
		id := rows[i].Id
		parentId := rows[i].Pid

		dp := make([]*DeptTree2, 0)
		mp[id] = &dp

		entry := &DeptTree2{Dept: &rows[i], Children: &dp}
		if parentId == pid {
			root = append(root, entry)
		} else {
			v, ok := mp[parentId]
			if ok {
				*v = append(*v, entry)
			} else {
				dp = []*DeptTree2{entry}
				mp[parentId] = &dp
			}
		}
	}
	return root
}
