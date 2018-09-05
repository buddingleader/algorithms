package insertionsort

// InsertionSort insertion sort
func InsertionSort(A []int, isIncrement bool) []int {
	for j := 1; j < len(A); j++ {
		i, key := j-1, A[j]
		for i >= 0 && A[i] > key && isIncrement {
			A[i+1] = A[i]
			i = i - 1
		}

		for i >= 0 && A[i] < key && !isIncrement {
			A[i+1] = A[i]
			i = i - 1
		}

		A[i+1] = key
	}
	return A
}
