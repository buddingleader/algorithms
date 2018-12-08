package heap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wangff15386/algorithms/common"
)

func Test_MaxHeapify(t *testing.T) {
	var (
		a  = []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
		a1 = []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	)
	MaxHeapify(a, len(a), 1)
	assert.Equal(t, a, a1)
}

func Test_BuildMaxHeap(t *testing.T) {
	var (
		a  = []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
		a1 = []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	)
	BuildMaxHeap(a, len(a))
	assert.Equal(t, a, a1)
}

func Test_MaxHeapSort(t *testing.T) {
	a := common.A
	fmt.Println("a:", a)
	MaxHeapSort(a)
	fmt.Println("a:", a)
}
