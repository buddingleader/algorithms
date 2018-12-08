package quicksort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wangff15386/algorithms/common"
)

var (
	A              = []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}
	ASortIncrement = []int{2, 4, 5, 6, 7, 8, 9, 11, 12, 13, 19, 21}
	ASortDecrement = []int{21, 19, 13, 12, 11, 9, 8, 7, 6, 5, 4, 2}

	B = []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
)

func Test_QuickSort(t *testing.T) {
	QuickSort(A, 0, len(A)-1, true)
	assert.Equal(t, A, ASortIncrement)

	QuickSort(A, 0, len(A)-1, false)
	assert.Equal(t, A, ASortDecrement)

	BSort := B
	QuickSort(B, 0, len(B)-1, true)
	assert.Equal(t, B, BSort)

	C := ASortIncrement
	QuickSort(C, 0, len(C)-1, true)
	assert.Equal(t, C, ASortIncrement)
}

func Benchmark_QuickSort(b *testing.B) {
	size := 1000
	A := make([]int, size, size)
	rand.Seed(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for index := 0; index < size; index++ {
			A[index] = rand.Intn(size)
		}

		b.StartTimer()
		QuickSort(A, 0, len(A)-1, true)
	}
}

func Test_RandomizedQuickSort(t *testing.T) {
	RandomizedQuickSort(A, 0, len(A)-1, true)
	assert.Equal(t, A, ASortIncrement)

	RandomizedQuickSort(A, 0, len(A)-1, false)
	assert.Equal(t, A, ASortDecrement)

	BSort := B
	RandomizedQuickSort(B, 0, len(B)-1, true)
	assert.Equal(t, B, BSort)

	C := ASortIncrement
	RandomizedQuickSort(C, 0, len(C)-1, true)
	assert.Equal(t, C, ASortIncrement)
}

func Benchmark_RandomizedQuickSort(b *testing.B) {
	size := 1000
	A := make([]int, size, size)
	rand.Seed(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for index := 0; index < size; index++ {
			A[index] = rand.Intn(size)
		}

		b.StartTimer()
		RandomizedQuickSort(A, 0, len(A)-1, true)
	}
}

func Test_reviewQuick(t *testing.T) {
	A := common.A
	reviewQuick(A, 0, len(A)-1)
	assert.Equal(t, A, common.ASortIncrement)
}
