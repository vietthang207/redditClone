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
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty")
	AssertNilNode(t, heap.Peek(), "Should be nil")

	heap = NewHeap(false)
	utils.AssertEqual(t, heap.Size(), 0, "Should start empty")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty")
	AssertNilNode(t, heap.Peek(), "Should be nil")

	heap = NewHeapWithCapacity(true, 10)
	utils.AssertEqual(t, heap.Size(), 0, "Should start empty")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty")
	AssertNilNode(t, heap.Peek(), "Should be nil")
}

func TestAddMin(t *testing.T) {
	heap = NewHeap(true)
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(1)
	utils.AssertEqual(t, 1, heap.Size(), "Should have size 1")
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(2)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(3)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")

	heap = NewHeap(true)
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(3)
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	add(2)
	AssertNodeEqual(t, 2, heap.Peek(), "Top value is wrong")
	add(1)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
}

func TestAddMax(t *testing.T) {
	heap = NewHeap(false)
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(1)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(2)
	AssertNodeEqual(t, 2, heap.Peek(), "Top value is wrong")
	add(3)
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")

	heap = NewHeap(false)
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(3)
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	add(2)
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	add(1)
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
}

func TestIsEmpty(t *testing.T) {
	heap = NewHeap(true)
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(1)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	utils.AssertFalse(t, heap.IsEmpty(), "Should not be empty.")
	add(2)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	utils.AssertFalse(t, heap.IsEmpty(), "Should not be empty.")

	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 1, top, "Polled top should be 1")
	} else {
		t.Fatal()
	}
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
	AssertNilNode(t, heap.Peek(), "Should be None.")
	add(207)
	AssertNodeEqual(t, 207, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 1, heap.Size(), "")
	add(208)
	AssertNodeEqual(t, 207, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 2, heap.Size(), "Should start empty.")
	add(209)
	AssertNodeEqual(t, 207, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 3, heap.Size(), "Should start empty.")

	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 207, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 2, heap.Size(), "Should start empty.")

	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 208, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 1, heap.Size(), "Should start empty.")

	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 209, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should start empty.")
}

func TestPollMin(t *testing.T) {
	heap = NewHeap(true)
	AssertNilNode(t, heap.Poll(), "Should be None.")
	add(10)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	add(11)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	add(12)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 11, top, "")
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
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	AssertNilNode(t, heap.Poll(), "Should be None.")

	//repeat the process again
	add(12)
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	add(11)
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	add(10)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 11, top, "")
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
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	AssertNilNode(t, heap.Poll(), "Should be None.")
}

func TestPollMax(t *testing.T) {
	heap = NewHeap(false)
	AssertNilNode(t, heap.Poll(), "Should be None.")
	add(10)
	AssertNodeEqual(t, 10, heap.Peek(), "Top value is wrong")
	add(11)
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	add(12)
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 11, top, "")
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
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
	AssertNilNode(t, heap.Poll(), "Should be None.")

	//repeat the process again
	add(12)
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	add(11)
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	add(10)
	AssertNodeEqual(t, 12, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 11, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 11, top, "")
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
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
}

func TestDeepHeapMin(t *testing.T) {
	heap = NewHeap(true)
	add(5)
	AssertNodeEqual(t, 5, heap.Peek(), "Top value is wrong")
	add(7)
	AssertNodeEqual(t, 5, heap.Peek(), "Top value is wrong")
	add(6)
	AssertNodeEqual(t, 5, heap.Peek(), "Top value is wrong")
	add(3)
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	add(4)
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	add(2)
	AssertNodeEqual(t, 2, heap.Peek(), "Top value is wrong")
	add(1)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(8)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(9)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(10)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	add(11)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 11, heap.Size(), "")
	add(12)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 12, heap.Size(), "")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 1, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 2, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 2, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 3, heap.Peek(), "Top value is wrong")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 3, top, "")
	} else {
		t.Fatal()
	}
	AssertNodeEqual(t, 4, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 9, heap.Size(), "")
	add(0)
	AssertNodeEqual(t, 0, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 10, heap.Size(), "")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 0, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 9, heap.Size(), "")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 4, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 8, heap.Size(), "")
	add(13)
	AssertNodeEqual(t, 5, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 9, heap.Size(), "")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 5, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 6, top, "")
	} else {
		t.Fatal()
	}
	add(1)
	AssertNodeEqual(t, 1, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 8, heap.Size(), "")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 1, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 7, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 8, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 9, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 10, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 11, top, "")
	} else {
		t.Fatal()
	}
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 12, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 1, heap.Size(), "")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 13, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
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
	add(-12)
	AssertNodeEqual(t, -1, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 12, heap.Size(), "")
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
	add(0)
	AssertNodeEqual(t, 0, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 10, heap.Size(), "")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, 0, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 9, heap.Size(), "")
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -4, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 8, heap.Size(), "")
	add(-13)
	AssertNodeEqual(t, -5, heap.Peek(), "Top value is wrong")
	utils.AssertEqual(t, 9, heap.Size(), "")
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
	if top := heap.Poll(); top != nil {
		AssertNodeEqual(t, -13, top, "")
	} else {
		t.Fatal()
	}
	utils.AssertEqual(t, 0, heap.Size(), "Should be empty.")
	utils.AssertTrue(t, heap.IsEmpty(), "Should be empty.")
	AssertNilNode(t, heap.Peek(), "Should be None.")
}
