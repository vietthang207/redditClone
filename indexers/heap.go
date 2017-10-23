package indexers

const (
	minHeap = true
	maxHeap = false
)

type Heap struct {
	isMin       bool
	nodes       []Node
	backPointer map[int]int
}

func NewHeap(isMin bool) *Heap {
	return NewHeapWithCapacity(isMin, 0)
}

func NewHeapWithCapacity(isMin bool, cap int) *Heap {
	h := new(Heap)
	h.isMin = isMin
	nodes := make([]Node, 0, cap+1)
	h.nodes = nodes
	backPointer := make(map[int]int, cap+1)
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
	h.backPointer[n.Id()] = len(h.nodes) - 1
	h.moveUp(h.Size() - 1)
}

func (h Heap) Peek() *Node {
	if h.IsEmpty() {
		return nil
	}
	tmp := h.nodes[0]
	return &Node{tmp.Key(), tmp.Id()}
}

func (h *Heap) Poll() *Node {
	if h.IsEmpty() {
		return nil
	}
	tmp := h.Peek()
	ret := &Node{tmp.Key(), tmp.Id()}
	delete(h.backPointer, ret.Id())
	if h.Size() == 1 {
		h.nodes = h.nodes[:0]
		return ret
	}
	h.nodes[0] = h.nodes[len(h.nodes)-1]
	h.backPointer[h.nodes[0].Id()] = 0
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.moveDown(0)
	return ret
}

func (h *Heap) Update(id int, key int) {
	pos := h.backPointer[id]
	h.nodes[pos].SetKey(key)
	h.moveUp(pos)
	h.moveDown(pos)
}

func (h *Heap) GetAllNodes() []Node {
	ret := make([]Node, h.Size())
	for i := 0; i < len(ret); i++ {
		ret[i] = *h.Poll()
	}
	for i := 0; i < len(ret); i++ {
		h.Add(ret[i])
	}
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
	backPointer := h.backPointer
	backPointer[nodes[pos1].Id()], backPointer[nodes[pos2].Id()] = pos2, pos1
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
