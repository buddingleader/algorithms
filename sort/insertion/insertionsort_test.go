package insertionsort

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

var (
	A              = []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}
	ASortIncrement = []int{2, 4, 5, 6, 7, 8, 9, 11, 12, 13, 19, 21}
	ASortDecrement = []int{21, 19, 13, 12, 11, 9, 8, 7, 6, 5, 4, 2}

	B = []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
)

func Test_InsertionSort(t *testing.T) {
	InsertionSort(A, true)
	assert.Equal(t, A, ASortIncrement)

	InsertionSort(A, false)
	assert.Equal(t, A, ASortDecrement)

	BSort := B
	InsertionSort(B, false)
	assert.Equal(t, B, BSort)
}
