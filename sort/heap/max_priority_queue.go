package heap

import "github.com/wangff15386/algorithms/common"

// Max return the root of heap
func Max(a []int) int {
	return a[0]
}

/* a is a max heap, so change the mind
// Insert x to a
func Insert(a []int, x int) []int {
	a = append(a, x)
	BuildMaxHeap(a, len(a))
	return a
}

// ExtractMax remove max
func ExtractMax(a []int) []int {
	heapsize := len(a) - 1
	common.Exchange(a, 0, heapsize)
	BuildMaxHeap(a, heapsize)
	return a[:heapsize]
}

// IncreaseKey increase the key of the index i
func IncreaseKey(a []int, i, key int) {
	a[i] = key
	BuildMaxHeap(a, len(a))
} */

// Insert x to a
func Insert(a []int, x int) []int {
	a = append(a, x)
	IncreaseKey(a, len(a)-1, x)
	return a
}

// ExtractMax remove max
func ExtractMax(a []int) ([]int, int) {
	max := a[0]
	a = a[1:]
	MaxHeapify(a, len(a), 0)
	return a, max
}

// IncreaseKey increase the key of the index i
func IncreaseKey(a []int, i, key int) {
	a[i] = key
	parent := (i - 1) / 2
	for i > 0 && a[i] > a[parent] {
		common.Exchange(a, i, parent)
		i = parent
		parent = (i - 1) / 2
	}
}
