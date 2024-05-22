package leet_03

import (
	"golang.org/x/exp/constraints"
)

type NodeSlice[T constraints.Ordered, E any] []Node[T, E]

func (nodes NodeSlice[T, E]) Len() int {
	return len(nodes)
}
func (nodes NodeSlice[T, E]) Less(i, j int) bool {
	return nodes[i].GetId() < nodes[j].GetId()
}
func (nodes NodeSlice[T, E]) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

type Node[T constraints.Ordered, E any] interface {
	GetId() T
	GetPid() T
	AppendChildren(E)
}

func IntoTree[T constraints.Ordered, E Node[T, E]](rows []E, rootPid T) []E {
	var root []E
	nodes := make(map[T]E, len(rows))
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

func (d *Dept) AppendChildren(v *Dept) {
	d.Children = append(d.Children, v)
}
