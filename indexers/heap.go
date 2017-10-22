package indexers

type Heap struct {
	isMin       bool
	nodes       []Node
	backPointer []int
}

func NewHeap(isMin bool) *Heap {
	return NewHeapWithCapacity(isMin, 0)
}

func NewHeapWithCapacity(isMin bool, cap int) *Heap {
	h := new(Heap)
	h.isMin = isMin
	nodes := make([]Node, 0, cap+1)
	h.nodes = nodes
	backPointer := make([]int, 0, cap+1)
	h.backPointer = backPointer
	return h
}

func (h Heap) Size() int {
	return len(h.nodes)
}

func (h Heap) IsEmpty() bool {
	return h.Size() < 1
}

func (h *Heap) Add(n Node) {
	h.nodes = append(h.nodes, n)
	h.moveUp(h.Size() - 1)
}

func (h Heap) Peek() *Node {
	if h.IsEmpty() {
		return nil
	}
	return &h.nodes[0]
}

func (h *Heap) Poll() *Node {
	if h.IsEmpty() {
		return nil
	}
	tmp := h.Peek()
	ret := &Node{tmp.Key(), tmp.Id()}
	if h.Size() == 1 {
		h.nodes = h.nodes[:0]
		return ret
	}
	h.nodes[0] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.moveDown(0)
	return ret
}

func (h Heap) leftChild(parent int) int {
	return (parent+1)*2 - 1
}
func (h Heap) rightChild(parent int) int {
	return (parent + 1) * 2
}
func (h Heap) parent(child int) int {
	return (child+1)/2 - 1
}
func (h Heap) isRoot(pos int) bool {
	return pos == 0
}
func (h Heap) isLeaf(pos int) bool {
	return h.leftChild(pos) >= h.Size()
}

func (h Heap) swap(pos1, pos2 int) {
	nodes := h.nodes
	// tmp := Node{nodes[pos1].Key(), nodes[pos1].Id()}
	// nodes[pos1] = nodes[pos2]
	// nodes[pos2] = tmp
	nodes[pos1], nodes[pos2] = nodes[pos2], nodes[pos1]
}

func (h Heap) moveUp(pos int) {
	nodes := h.nodes
	parentPos := h.parent(pos)
	if h.isMin {
		for !h.isRoot(pos) && nodes[pos].Less(nodes[parentPos]) {
			h.swap(pos, parentPos)
			pos = parentPos
			parentPos = h.parent(pos)
		}
	} else {
		for !h.isRoot(pos) && !nodes[pos].Less(nodes[parentPos]) {
			h.swap(pos, parentPos)
			pos = parentPos
			parentPos = h.parent(pos)
		}
	}
}

func (h Heap) moveDown(pos int) {
	nodes := h.nodes
	if h.isMin {
		for !h.isLeaf(pos) {
			smallChildPos := h.leftChild(pos)
			if smallChildPos >= h.Size() {
				break
			}
			if h.rightChild(pos) < h.Size() && nodes[h.rightChild(pos)].Less(nodes[smallChildPos]) {
				smallChildPos = h.rightChild(pos)
			}
			if nodes[smallChildPos].Less(nodes[pos]) {
				h.swap(pos, smallChildPos)
				pos = smallChildPos
			} else {
				break
			}
		}
	} else {
		for !h.isLeaf(pos) {
			bigChildPos := h.leftChild(pos)
			if bigChildPos >= h.Size() {
				break
			}
			if h.rightChild(pos) < h.Size() && !nodes[h.rightChild(pos)].Less(nodes[bigChildPos]) {
				bigChildPos = h.rightChild(pos)
			}
			if !nodes[bigChildPos].Less(nodes[pos]) {
				h.swap(pos, bigChildPos)
				pos = bigChildPos
			} else {
				break
			}
		}
	}
}
