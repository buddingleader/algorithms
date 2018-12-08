package heap

import (
	"github.com/wangff15386/algorithms/common"
)

// MaxHeapify keep the max heap
func MaxHeapify(a []int, heapsize, i int) {
	l, r := 2*i+1, 2*i+2
	largest := i
	if l < heapsize && a[l] > a[i] {
		largest = l
	}

	if r < heapsize && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		common.Exchange(a, i, largest)
		MaxHeapify(a, heapsize, largest)
	}
}

// BuildMaxHeap build a max heap from the last no leaf node to root node
func BuildMaxHeap(a []int, heapsize int) {
	for index := len(a) / 2; index >= 0; index-- {
		MaxHeapify(a, heapsize, index)
	}
}

// MaxHeapSort use max heap to sort a array, increament
func MaxHeapSort(a []int) {
	heapsize := len(a)
	for index := 0; index < len(a); index++ {
		BuildMaxHeap(a, heapsize)
		heapsize--
		common.Exchange(a, 0, heapsize)
	}
}
