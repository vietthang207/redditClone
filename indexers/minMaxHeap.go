package indexers

const (
	QueryMin    = true
	QueryMax    = false
	inQueryHeap = true
	inBigHeap   = false
)

type MinMaxHeap struct {
	isQueryMin    bool
	querySize     int
	queryHeap     *Heap
	bigHeap       *Heap
	heapLocations map[int]bool
}

func NewMinMaxHeap(isQueryMin bool, querySize int) *MinMaxHeap {
	mmHeap := new(MinMaxHeap)
	mmHeap.isQueryMin = isQueryMin
	mmHeap.querySize = querySize
	mmHeap.queryHeap = NewHeap(!isQueryMin)
	mmHeap.bigHeap = NewHeap(isQueryMin)
	mmHeap.heapLocations = make(map[int]bool)
	return mmHeap
}

func (h *MinMaxHeap) Size() int {
	return h.queryHeap.Size() + h.bigHeap.Size()
}

func (h *MinMaxHeap) AddById(id int, key int) {
	h.Add(Node{key, id})
}

func (h *MinMaxHeap) Add(n Node) {
	if h.queryHeap.Size() < h.querySize {
		h.queryHeap.Add(n)
		h.heapLocations[n.Id()] = inQueryHeap
		return
	}
	if h.isQueryMin {
		if n.Key() < h.queryHeap.Peek().Key() {
			tmp := *h.queryHeap.Poll()
			h.heapLocations[tmp.Id()] = inBigHeap
			h.bigHeap.Add(tmp)
			h.queryHeap.Add(n)
			h.heapLocations[n.Id()] = inQueryHeap
		} else {
			h.bigHeap.Add(n)
			h.heapLocations[n.Id()] = inBigHeap
		}
	} else {
		if n.Key() > h.queryHeap.Peek().Key() {
			tmp := *h.queryHeap.Poll()
			h.heapLocations[tmp.Id()] = inBigHeap
			h.bigHeap.Add(tmp)
			h.queryHeap.Add(n)
			h.heapLocations[n.Id()] = inQueryHeap
		} else {
			h.bigHeap.Add(n)
			h.heapLocations[n.Id()] = inBigHeap
		}
	}
}

func (h *MinMaxHeap) Update(id int, key int) {
	if h.heapLocations[id] == inQueryHeap {
		h.queryHeap.Update(id, key)
	} else {
		h.bigHeap.Update(id, key)
	}
	h.checkTopOfTwoHeaps()
}

func (h *MinMaxHeap) checkTopOfTwoHeaps() {
	if (h.isQueryMin && h.bigHeap.Peek().Key() < h.queryHeap.Peek().Key()) || (!h.isQueryMin && h.bigHeap.Peek().Key() > h.queryHeap.Peek().Key()) {
		tmp1 := *h.queryHeap.Poll()
		tmp2 := *h.bigHeap.Poll()
		h.heapLocations[tmp1.Id()] = inBigHeap
		h.heapLocations[tmp2.Id()] = inQueryHeap
		h.queryHeap.Add(tmp2)
		h.bigHeap.Add(tmp1)
		h.queryHeap.moveDown(0)
		h.bigHeap.moveDown(0)
	}
}

func (h *MinMaxHeap) Query() []Node {
	tmp := h.queryHeap.GetAllNodes()
	ret := make([]Node, len(tmp))
	for i := 0; i < len(tmp); i++ {
		ret[i] = tmp[len(tmp)-1-i]
	}
	return ret
}
