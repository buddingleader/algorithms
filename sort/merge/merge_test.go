package merge

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wangff15386/algorithms/common"
)

func Test_MergeSort(t *testing.T) {
	A := common.A
	Sort(A, 0, len(A)-1)
	assert.Equal(t, A, common.ASortIncrement)
}
