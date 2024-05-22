package leet_03

type Node interface {
	GetId() int
	GetPid() int
	AppendChildren(Node)
}

func IntoTree[T Node](rows []T, rootPid int) []T {
	var root []T
	nodes := make(map[int]T, len(rows))
	for _, v := range rows {
		node := v
		id := node.GetId()
		pid := node.GetPid()
		if pid == rootPid {
			root = append(root, node)
		} else if parent, exists := nodes[pid]; exists {
			parent.AppendChildren(node)
		}
		nodes[id] = node
	}
	return root
}

type Dept struct {
	Id       int
	Pid      int
	Name     string
	Children []*Dept
}

func (d *Dept) GetId() int {
	return d.Id
}

func (d *Dept) GetPid() int {
	return d.Pid
}

func (d *Dept) AppendChildren(v Node) {
	d.Children = append(d.Children, v.(*Dept))
}
