package quicksort

import (
	"math/rand"
)

// QuickSort quick sort
func QuickSort(A []int, start, end int, isIncrement bool) []int {
	if start < end {
		r := partition(A, start, end, isIncrement)
		QuickSort(A, start, r-1, isIncrement)
		QuickSort(A, r+1, end, isIncrement)
	}

	return A
}

// Partition A with A[end], less than A[end] on the left, more than A[end] on the right
// When partitioning, A is divided into three blocks, [start, p] is less than A [end], [p + 1, q] is more than A [end], [q, end-1] is waiting for partition
func partition(A []int, start, end int, isIncrement bool) int {
	p, x := start-1, end
	for q := start; q <= end; q++ {
		if A[x] > A[q] && isIncrement {
			p++
			exchange(A, p, q)
		}

		if A[x] < A[q] && !isIncrement {
			p++
			exchange(A, p, q)
		}
	}

	exchange(A, p+1, x)
	return p + 1
}

func exchange(A []int, x, y int) {
	temp := A[x]
	A[x] = A[y]
	A[y] = temp
}

// RandomizedQuickSort quick sort for randomized x
func RandomizedQuickSort(A []int, start, end int, isIncrement bool) []int {
	if start < end {
		r := randomizedPartition(A, start, end, isIncrement)
		RandomizedQuickSort(A, start, r-1, isIncrement)
		RandomizedQuickSort(A, r+1, end, isIncrement)
	}

	return A
}

// RandomizedPartition A with a randomize A[end], less than A[end] on the left, more than A[end] on the right
// The A[end] is randomly selected by another element exchanged with A[end]
func randomizedPartition(A []int, start, end int, isIncrement bool) int {
	// If you set a random number seed for each partition, the performance is too low.
	// rand.Seed(time.Now().Unix())
	index := rand.Intn(end-start) + start
	exchange(A, end, index)
	return partition(A, start, end, isIncrement)
}
