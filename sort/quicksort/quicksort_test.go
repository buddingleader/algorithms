package quicksort

import (
	"math/rand"
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_QuickSort(t *testing.T) {
	A := []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}
	ASortIncrement := []int{2, 4, 5, 6, 7, 8, 9, 11, 12, 13, 19, 21}
	QuickSort(A, 0, len(A)-1, true)
	assert.Equal(t, A, ASortIncrement)

	ASortDecrement := []int{21, 19, 13, 12, 11, 9, 8, 7, 6, 5, 4, 2}
	QuickSort(A, 0, len(A)-1, false)
	assert.Equal(t, A, ASortDecrement)

	B := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	BSort := B
	QuickSort(B, 0, len(B)-1, true)
	assert.Equal(t, B, BSort)
}

func Benchmark_QuickSort(b *testing.B) {
	size := 1000
	A := make([]int, size, size)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for index := 0; index < size; index++ {
			A[index] = rand.Intn(size)
		}

		b.StartTimer()
		QuickSort(A, 0, len(A)-1, true)
	}

}
