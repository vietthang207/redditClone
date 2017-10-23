package indexers

import (
	"redditClone/utils"
	"testing"
)

var (
	heap = NewHeap(true)
)

func add(i int) {
	heap.Add(Node{i, i})
}

func AssertNilNode(t *testing.T, a *Node, message string) {
	if a != nil {
		t.Errorf(message)
	}
}

func AssertNodeEqual(t *testing.T, i int, a *Node, message string) {
	utils.AssertEqual(t, Node{i, i}, *a, message)
}
func TestNewHeap(t *testing.T) {
	heap = NewHeap(true)
	utils.AssertEqual(t, heap.Size(), 0, "Should start empty")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty")
	AssertNilNode(t, heap.Peek(), "Should be nil")

	heap = NewHeap(false)
	utils.AssertEqual(t, heap.Size(), 0, "Should start empty")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty")
	AssertNilNode(t, heap.Peek(), "Should be nil")

	heap = NewHeapWithCapacity(true, 10)
	utils.AssertEqual(t, heap.Size(), 0, "Should start empty")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty")
	AssertNilNode(t, heap.Peek(), "Should be nil")
}

func TestAddMin(t *testing.T) {
	heap = NewHeap(true)
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(1)
	utils.AssertEqual(t, 1, heap.Size(), "Should have size 1")
	utils.AssertEqual(t, 1, len(heap.backPointer), "Should start empty.")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0}, "Backpointer is wrong")
	add(2)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 2: 1}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(3)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 2: 1, 3: 2}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")

	heap = NewHeap(true)
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(3)
	utils.AssertMap(t, heap.backPointer, map[int]int{3: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	add(2)
	utils.AssertMap(t, heap.backPointer, map[int]int{2: 0, 3: 1}, "Backpointer is wrong")
	AssertNodeEqual(t, 2, heap.Peek(), "Top value is wrong")
	add(1)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 2: 2, 3: 1}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
}

func TestAddMax(t *testing.T) {
	heap = NewHeap(false)
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(1)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(2)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 1, 2: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 2, heap.Peek(), "Top value is wrong")
	add(3)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 1, 2: 2, 3: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")

	heap = NewHeap(false)
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(3)
	utils.AssertMap(t, heap.backPointer, map[int]int{3: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	add(2)
	utils.AssertMap(t, heap.backPointer, map[int]int{2: 1, 3: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	add(1)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 2, 2: 1, 3: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
}

func TestIsEmpty(t *testing.T) {
	heap = NewHeap(true)
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(1)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	utils.AssertFalse(t, heap.IsEmpty(), "Should not be empty.")
	add(2)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 2: 1}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	utils.AssertFalse(t, heap.IsEmpty(), "Should not be empty.")

	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 1, top, "Polled top should be 1")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{2: 0}, "Backpointer is wrong")
	utils.AssertFalse(t, heap.IsEmpty(), "Should not be empty.")

	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 2, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
}

func TestSize(t *testing.T) {
	heap = NewHeap(true)
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(207)
	utils.AssertMap(t, heap.backPointer, map[int]int{207: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 207, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 1, len(heap.backPointer), "Should start empty.")
	utils.AssertEqual(t, 1, heap.Size(), "")
	add(208)
	utils.AssertMap(t, heap.backPointer, map[int]int{207: 0, 208: 1}, "Backpointer is wrong")
	AssertNodeEqual(t, 207, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 2, heap.Size(), "Should start empty.")
	utils.AssertEqual(t, 2, len(heap.backPointer), "Should start empty.")
	add(209)
	utils.AssertMap(t, heap.backPointer, map[int]int{207: 0, 208: 1, 209: 2}, "Backpointer is wrong")
	AssertNodeEqual(t, 207, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 3, len(heap.backPointer), "Should start empty.")
	utils.AssertEqual(t, 3, heap.Size(), "Should start empty.")

	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 207, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{208: 0, 209: 1}, "Backpointer is wrong")
	utils.AssertEqual(t, 2, heap.Size(), "Should start empty.")
	utils.AssertEqual(t, 2, len(heap.backPointer), "Should start empty.")

	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 208, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 1, heap.Size(), "Should start empty.")
	utils.AssertMap(t, heap.backPointer, map[int]int{209: 0}, "Backpointer is wrong")
	utils.AssertEqual(t, 1, len(heap.backPointer), "Should start empty.")

	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 209, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertMap(t, heap.backPointer, map[int]int{}, "Backpointer is wrong")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
}

func TestPollMin(t *testing.T) {
	heap = NewHeap(true)
	AssertNilNode(t, heap.Poll(), "Should be None.")
	add(10)
	utils.AssertMap(t, heap.backPointer, map[int]int{10: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	add(11)
	utils.AssertMap(t, heap.backPointer, map[int]int{10: 0, 11: 1}, "Backpointer is wrong")
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	add(12)
	utils.AssertMap(t, heap.backPointer, map[int]int{10: 0, 11: 1, 12: 2}, "Backpointer is wrong")
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	utils.AssertMap(t, heap.backPointer, map[int]int{11: 0, 12: 1}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 11, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	utils.AssertMap(t, heap.backPointer, map[int]int{12: 0}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
	utils.AssertMap(t, heap.backPointer, map[int]int{}, "Backpointer is wrong")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	AssertNilNode(t, heap.Poll(), "Should be None.")

	//repeat the process again
	add(12)
	utils.AssertMap(t, heap.backPointer, map[int]int{12: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	add(11)
	utils.AssertMap(t, heap.backPointer, map[int]int{12: 1, 11: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	add(10)
	utils.AssertMap(t, heap.backPointer, map[int]int{12: 1, 11: 2, 10: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{12: 1, 11: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 11, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	utils.AssertMap(t, heap.backPointer, map[int]int{12: 0}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
	utils.AssertMap(t, heap.backPointer, map[int]int{}, "Backpointer is wrong")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	AssertNilNode(t, heap.Poll(), "Should be None.")
}

func TestPollMax(t *testing.T) {
	heap = NewHeap(false)
	AssertNilNode(t, heap.Poll(), "Should be None.")
	add(10)
	utils.AssertMap(t, heap.backPointer, map[int]int{10: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	add(11)
	utils.AssertMap(t, heap.backPointer, map[int]int{11: 0, 10: 1}, "Backpointer is wrong")
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	add(12)
	utils.AssertMap(t, heap.backPointer, map[int]int{12: 0, 10: 1, 11: 2}, "Backpointer is wrong")
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{11: 0, 10: 1}, "Backpointer is wrong")
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 11, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	utils.AssertMap(t, heap.backPointer, map[int]int{10: 0}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
	utils.AssertMap(t, heap.backPointer, map[int]int{}, "Backpointer is wrong")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	AssertNilNode(t, heap.Poll(), "Should be None.")

	//repeat the process again
	add(12)
	utils.AssertMap(t, heap.backPointer, map[int]int{12: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	add(11)
	utils.AssertMap(t, heap.backPointer, map[int]int{12: 0, 11: 1}, "Backpointer is wrong")
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	add(10)
	utils.AssertMap(t, heap.backPointer, map[int]int{12: 0, 11: 1, 10: 2}, "Backpointer is wrong")
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	utils.AssertMap(t, heap.backPointer, map[int]int{11: 0, 10: 1}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 11, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	utils.AssertMap(t, heap.backPointer, map[int]int{10: 0}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
	utils.AssertMap(t, heap.backPointer, map[int]int{}, "Backpointer is wrong")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	AssertNilNode(t, heap.Poll(), "Should be None.")
}

func TestDuplicateMin(t *testing.T) {
	heap = NewHeap(true)
	add(10)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	add(10)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	add(12)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")

	add(12)
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	add(10)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	add(10)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
}

func TestDuplicateMax(t *testing.T) {
	heap = NewHeap(false)
	add(10)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	add(10)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	add(12)
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")

	add(12)
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	add(10)
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	add(10)
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
	utils.AssertEqual(t, 0, len(heap.backPointer), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
}

func TestDeepHeapMin(t *testing.T) {
	heap = NewHeap(true)
	add(5)
	utils.AssertMap(t, heap.backPointer, map[int]int{5: 0}, "Backpointer is wrong")
	AssertNodeEqual(t, 5, heap.Peek(), "Top value is wrong")
	add(7)
	utils.AssertMap(t, heap.backPointer, map[int]int{5: 0, 7: 1}, "Backpointer is wrong")
	AssertNodeEqual(t, 5, heap.Peek(), "Top value is wrong")
	add(6)
	utils.AssertMap(t, heap.backPointer, map[int]int{5: 0, 7: 1, 6: 2}, "Backpointer is wrong")
	AssertNodeEqual(t, 5, heap.Peek(), "Top value is wrong")
	add(3)
	utils.AssertMap(t, heap.backPointer, map[int]int{3: 0, 7: 3, 6: 2, 5: 1}, "Backpointer is wrong")
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	add(4)
	utils.AssertMap(t, heap.backPointer, map[int]int{3: 0, 4: 1, 7: 3, 6: 2, 5: 4}, "Backpointer is wrong")
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	add(2)
	utils.AssertMap(t, heap.backPointer, map[int]int{2: 0, 3: 2, 4: 1, 7: 3, 6: 5, 5: 4}, "Backpointer is wrong")
	AssertNodeEqual(t, 2, heap.Peek(), "Top value is wrong")
	add(1)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 2: 2, 3: 6, 4: 1, 7: 3, 6: 5, 5: 4}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(8)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 2: 2, 3: 6, 4: 1, 7: 3, 6: 5, 5: 4, 8: 7}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(9)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 2: 2, 3: 6, 4: 1, 7: 3, 6: 5, 5: 4, 8: 7, 9: 8}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(10)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 2: 2, 3: 6, 4: 1, 7: 3, 6: 5, 5: 4, 8: 7, 9: 8, 10: 9}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(11)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 2: 2, 3: 6, 4: 1, 7: 3, 6: 5, 5: 4, 8: 7, 9: 8, 10: 9, 11: 10}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 11, heap.Size(), "")
	utils.AssertEqual(t, 11, len(heap.backPointer), "Should start empty.")
	add(12)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 2: 2, 3: 6, 4: 1, 7: 3, 6: 5, 5: 4, 8: 7, 9: 8, 10: 9, 11: 10, 12: 11}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 12, heap.Size(), "")
	utils.AssertEqual(t, 12, len(heap.backPointer), "Should start empty.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 1, top, "")
	} else {
		t.Fatal()
		utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 2: 2, 3: 6, 4: 1, 7: 3, 6: 5, 5: 4, 8: 7, 9: 8, 10: 9, 11: 10, 12: 11}, "Backpointer is wrong")
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{2: 0, 3: 2, 4: 1, 7: 3, 6: 5, 5: 4, 8: 7, 9: 8, 10: 9, 11: 10, 12: 6}, "Backpointer is wrong")
	AssertNodeEqual(t, 2, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 2, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{3: 0, 4: 1, 7: 3, 6: 2, 5: 4, 8: 7, 9: 8, 10: 9, 11: 5, 12: 6}, "Backpointer is wrong")
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 3, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{4: 0, 7: 3, 6: 2, 5: 1, 8: 7, 9: 8, 10: 4, 11: 5, 12: 6}, "Backpointer is wrong")
	AssertNodeEqual(t, 4, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 9, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	add(0)
	utils.AssertMap(t, heap.backPointer, map[int]int{0: 0, 4: 1, 7: 3, 6: 2, 5: 4, 8: 7, 9: 8, 10: 9, 11: 5, 12: 6}, "Backpointer is wrong")
	AssertNodeEqual(t, 0, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 10, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 0, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{4: 0, 7: 3, 6: 2, 5: 1, 8: 7, 9: 8, 10: 4, 11: 5, 12: 6}, "Backpointer is wrong")
	utils.AssertEqual(t, 9, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 4, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{7: 1, 6: 2, 5: 0, 8: 3, 9: 7, 10: 4, 11: 5, 12: 6}, "Backpointer is wrong")
	utils.AssertEqual(t, 8, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	add(13)
	utils.AssertMap(t, heap.backPointer, map[int]int{7: 1, 6: 2, 5: 0, 8: 3, 9: 7, 10: 4, 11: 5, 12: 6, 13: 8}, "Backpointer is wrong")
	AssertNodeEqual(t, 5, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 9, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 5, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{7: 1, 6: 0, 8: 3, 9: 7, 10: 4, 11: 2, 12: 6, 13: 5}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 6, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{7: 0, 8: 1, 9: 3, 10: 4, 11: 2, 12: 6, 13: 5}, "Backpointer is wrong")
	add(1)
	utils.AssertMap(t, heap.backPointer, map[int]int{1: 0, 7: 1, 8: 3, 9: 7, 10: 4, 11: 2, 12: 6, 13: 5}, "Backpointer is wrong")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 8, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 1, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{7: 0, 8: 1, 9: 3, 10: 4, 11: 2, 12: 6, 13: 5}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 7, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{8: 0, 9: 1, 10: 4, 11: 2, 12: 3, 13: 5}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 8, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{9: 0, 10: 1, 11: 2, 12: 3, 13: 4}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 9, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{10: 0, 11: 2, 12: 1, 13: 3}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{11: 0, 12: 1, 13: 2}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 11, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{12: 0, 13: 1}, "Backpointer is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{13: 0}, "Backpointer is wrong")
	utils.AssertEqual(t, 1, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 13, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertMap(t, heap.backPointer, map[int]int{}, "Backpointer is wrong")
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
}
func TestDeepHeapMax(t *testing.T) {
	heap = NewHeap(false)
	add(-5)
	AssertNodeEqual(t, -5, heap.Peek(), "Top value is wrong")
	add(-7)
	AssertNodeEqual(t, -5, heap.Peek(), "Top value is wrong")
	add(-6)
	AssertNodeEqual(t, -5, heap.Peek(), "Top value is wrong")
	add(-3)
	AssertNodeEqual(t, -3, heap.Peek(), "Top value is wrong")
	add(-4)
	AssertNodeEqual(t, -3, heap.Peek(), "Top value is wrong")
	add(-2)
	AssertNodeEqual(t, -2, heap.Peek(), "Top value is wrong")
	add(-1)
	AssertNodeEqual(t, -1, heap.Peek(), "Top value is wrong")
	add(-8)
	AssertNodeEqual(t, -1, heap.Peek(), "Top value is wrong")
	add(-9)
	AssertNodeEqual(t, -1, heap.Peek(), "Top value is wrong")
	add(-10)
	AssertNodeEqual(t, -1, heap.Peek(), "Top value is wrong")
	add(-11)
	AssertNodeEqual(t, -1, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 11, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	add(-12)
	AssertNodeEqual(t, -1, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 12, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -1, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, -2, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -2, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, -3, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -3, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, -4, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 9, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	add(0)
	AssertNodeEqual(t, 0, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 10, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 0, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 9, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -4, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 8, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	add(-13)
	AssertNodeEqual(t, -5, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 9, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -5, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -6, top, "")
	} else {
		t.Fatal()
	}
	add(-1)
	AssertNodeEqual(t, -1, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 8, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -1, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -7, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -8, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -9, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -10, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -11, top, "This should be -11")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -12, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 1, heap.Size(), "")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -13, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
	utils.AssertEqual(t, heap.Size(), len(heap.backPointer), "Backpointer and nodes length does not match.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
}
