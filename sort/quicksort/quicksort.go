package quicksort

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
	x := end
	p := start - 1
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
