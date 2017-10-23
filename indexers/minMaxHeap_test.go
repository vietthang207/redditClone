package indexers

import "testing"
import "redditClone/utils"

var (
	h = NewMinMaxHeap(QueryMin, 4)
)

func addId(id int, key int) {
	h.Add(Node{id: id, key: key})
}

func AssertIdList(t *testing.T, a []Node, b []int, message string) {
	list := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		list[i] = a[i].Id()
	}
	utils.AssertList(t, list, b, message)
}

func TestNewMinMaxHeap(t *testing.T) {
	h = NewMinMaxHeap(QueryMin, 4)
	utils.AssertEqual(t, h.querySize, 4, "")
}

func TestMinMamHeapQueryMax(t *testing.T) {
	h = NewMinMaxHeap(QueryMax, 4)
	addId(1, 1)
	utils.AssertEqual(t, h.queryHeap.Size(), 1, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 0, "")
	AssertIdList(t, h.Query(), []int{1}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{1}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{}, "")

	addId(2, 2)
	utils.AssertEqual(t, h.queryHeap.Size(), 2, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 0, "")
	AssertIdList(t, h.Query(), []int{2, 1}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{1, 2}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{}, "")

	addId(3, 3)
	utils.AssertEqual(t, h.queryHeap.Size(), 3, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 0, "")
	AssertIdList(t, h.Query(), []int{3, 2, 1}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{1, 2, 3}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{}, "")

	addId(4, 4)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 0, "")
	AssertIdList(t, h.Query(), []int{4, 3, 2, 1}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{1, 2, 3, 4}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{}, "")

	addId(5, 5)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 1, "")
	AssertIdList(t, h.Query(), []int{5, 4, 3, 2}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{2, 3, 4, 5}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{1}, "")

	addId(6, 6)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 2, "")
	AssertIdList(t, h.Query(), []int{6, 5, 4, 3}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{3, 4, 5, 6}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{2, 1}, "")

	addId(0, 0)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 3, "")
	AssertIdList(t, h.Query(), []int{6, 5, 4, 3}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{3, 4, 5, 6}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{2, 1, 0}, "")

	addId(7, 7)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 4, "")
	AssertIdList(t, h.Query(), []int{7, 6, 5, 4}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{4, 5, 6, 7}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{3, 2, 1, 0}, "")

	addId(8, 8)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 5, "")
	AssertIdList(t, h.Query(), []int{8, 7, 6, 5}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{5, 6, 7, 8}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{4, 3, 2, 1, 0}, "")

	addId(9, 9)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 6, "")
	AssertIdList(t, h.Query(), []int{9, 8, 7, 6}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{6, 7, 8, 9}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{5, 4, 3, 2, 1, 0}, "")

	addId(10, 10)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 7, "")
	AssertIdList(t, h.Query(), []int{10, 9, 8, 7}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{7, 8, 9, 10}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{6, 5, 4, 3, 2, 1, 0}, "")

	addId(11, 11)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 8, "")
	utils.AssertEqual(t, h.Size(), 12, "")
	AssertIdList(t, h.Query(), []int{11, 10, 9, 8}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{8, 9, 10, 11}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{7, 6, 5, 4, 3, 2, 1, 0}, "")

	h.Update(0, 20)
	AssertIdList(t, h.Query(), []int{0, 11, 10, 9}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{9, 10, 11, 0}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{8, 7, 6, 5, 4, 3, 2, 1}, "")

	h.Update(10, 12)
	AssertIdList(t, h.Query(), []int{0, 10, 11, 9}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{9, 11, 10, 0}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{8, 7, 6, 5, 4, 3, 2, 1}, "")

	h.Update(0, 0)
	AssertIdList(t, h.Query(), []int{10, 11, 9, 8}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{8, 9, 11, 10}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{7, 6, 5, 4, 3, 2, 1, 0}, "")

	h.Update(4, -10)
	AssertIdList(t, h.Query(), []int{10, 11, 9, 8}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{8, 9, 11, 10}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{7, 6, 5, 3, 2, 1, 0, 4}, "")

	h.Update(6, 4)
	AssertIdList(t, h.Query(), []int{10, 11, 9, 8}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{8, 9, 11, 10}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{7, 5, 6, 3, 2, 1, 0, 4}, "")

	h.Update(2, 20)
	AssertIdList(t, h.Query(), []int{2, 10, 11, 9}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{9, 11, 10, 2}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{8, 7, 5, 6, 3, 1, 0, 4}, "")

	h.Update(6, 19)
	AssertIdList(t, h.Query(), []int{2, 6, 10, 11}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{11, 10, 6, 2}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{9, 8, 7, 5, 3, 1, 0, 4}, "")

	h.Update(10, 6)
	AssertIdList(t, h.Query(), []int{2, 6, 11, 9}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{9, 11, 6, 2}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{8, 7, 10, 5, 3, 1, 0, 4}, "")
}

func TestMinMamHeapQueryMin(t *testing.T) {
	h = NewMinMaxHeap(QueryMin, 4)
	addId(1, -1)
	utils.AssertEqual(t, h.queryHeap.Size(), 1, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 0, "")
	AssertIdList(t, h.Query(), []int{1}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{1}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{}, "")

	addId(2, -2)
	utils.AssertEqual(t, h.queryHeap.Size(), 2, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 0, "")
	AssertIdList(t, h.Query(), []int{2, 1}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{1, 2}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{}, "")

	addId(3, -3)
	utils.AssertEqual(t, h.queryHeap.Size(), 3, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 0, "")
	AssertIdList(t, h.Query(), []int{3, 2, 1}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{1, 2, 3}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{}, "")

	addId(4, -4)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 0, "")
	AssertIdList(t, h.Query(), []int{4, 3, 2, 1}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{1, 2, 3, 4}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{}, "")

	addId(5, -5)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 1, "")
	AssertIdList(t, h.Query(), []int{5, 4, 3, 2}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{2, 3, 4, 5}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{1}, "")

	addId(6, -6)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 2, "")
	AssertIdList(t, h.Query(), []int{6, 5, 4, 3}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{3, 4, 5, 6}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{2, 1}, "")

	addId(0, 0)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 3, "")
	AssertIdList(t, h.Query(), []int{6, 5, 4, 3}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{3, 4, 5, 6}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{2, 1, 0}, "")

	addId(7, -7)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 4, "")
	AssertIdList(t, h.Query(), []int{7, 6, 5, 4}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{4, 5, 6, 7}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{3, 2, 1, 0}, "")

	addId(8, -8)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 5, "")
	AssertIdList(t, h.Query(), []int{8, 7, 6, 5}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{5, 6, 7, 8}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{4, 3, 2, 1, 0}, "")

	addId(9, -9)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 6, "")
	AssertIdList(t, h.Query(), []int{9, 8, 7, 6}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{6, 7, 8, 9}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{5, 4, 3, 2, 1, 0}, "")

	addId(10, -10)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 7, "")
	AssertIdList(t, h.Query(), []int{10, 9, 8, 7}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{7, 8, 9, 10}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{6, 5, 4, 3, 2, 1, 0}, "")

	addId(11, -11)
	utils.AssertEqual(t, h.queryHeap.Size(), 4, "")
	utils.AssertEqual(t, h.bigHeap.Size(), 8, "")
	utils.AssertEqual(t, h.Size(), 12, "")
	AssertIdList(t, h.Query(), []int{11, 10, 9, 8}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{8, 9, 10, 11}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{7, 6, 5, 4, 3, 2, 1, 0}, "")

	h.Update(0, -20)
	AssertIdList(t, h.Query(), []int{0, 11, 10, 9}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{9, 10, 11, 0}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{8, 7, 6, 5, 4, 3, 2, 1}, "")

	h.Update(10, -12)
	AssertIdList(t, h.Query(), []int{0, 10, 11, 9}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{9, 11, 10, 0}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{8, 7, 6, 5, 4, 3, 2, 1}, "")

	h.Update(0, 0)
	AssertIdList(t, h.Query(), []int{10, 11, 9, 8}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{8, 9, 11, 10}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{7, 6, 5, 4, 3, 2, 1, 0}, "")

	h.Update(4, 10)
	AssertIdList(t, h.Query(), []int{10, 11, 9, 8}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{8, 9, 11, 10}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{7, 6, 5, 3, 2, 1, 0, 4}, "")

	h.Update(6, -4)
	AssertIdList(t, h.Query(), []int{10, 11, 9, 8}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{8, 9, 11, 10}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{7, 5, 6, 3, 2, 1, 0, 4}, "")

	h.Update(2, -20)
	AssertIdList(t, h.Query(), []int{2, 10, 11, 9}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{9, 11, 10, 2}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{8, 7, 5, 6, 3, 1, 0, 4}, "")

	h.Update(6, -19)
	AssertIdList(t, h.Query(), []int{2, 6, 10, 11}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{11, 10, 6, 2}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{9, 8, 7, 5, 3, 1, 0, 4}, "")

	h.Update(10, -6)
	AssertIdList(t, h.Query(), []int{2, 6, 11, 9}, "")
	AssertIdList(t, h.queryHeap.GetAllNodes(), []int{9, 11, 6, 2}, "")
	AssertIdList(t, h.bigHeap.GetAllNodes(), []int{8, 7, 10, 5, 3, 1, 0, 4}, "")
}
