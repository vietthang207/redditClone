package indexers

import "errors"

var (
	ErrorNegativeCapacity = errors.New("Capacity need to be non-negative")
	ErrorIdNotFound       = errors.New("Id not found")
)

type Comparable interface {
	Less(Comparable) bool
}

type Node struct {
	key int
	id  int
}

func (n Node) Key() int {
	return n.key
}

func (n *Node) SetKey(key int) {
	n.key = key
}

func (n Node) Id() int {
	return n.id
}

func (n Node) Less(o Comparable) bool {
	otherNode, _ := o.(Node)
	return n.key < otherNode.key
}
