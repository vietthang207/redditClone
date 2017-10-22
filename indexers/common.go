package indexers

import "errors"

var (
	ErrorNegativeCapacity = errors.New("Capacity need to be non-negative")
)

type Node struct {
	key int
	id  int
}

func (n Node) Key() int {
	return n.key
}

func (n Node) Id() int {
	return n.id
}

func (n Node) Less(o Node) bool {
	return n.key < o.key
}
