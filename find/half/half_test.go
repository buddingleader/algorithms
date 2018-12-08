package half

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wangff15386/algorithms/common"
)

func Test_Find(t *testing.T) {
	A := common.ASortIncrement
	index, ok := Find(A, 0, len(A)-1, 5)
	assert.Equal(t, index, 2)
	assert.Equal(t, ok, true)

	index, ok = Find(A, 0, len(A)-1, 20)
	assert.Equal(t, index, -1)
	assert.Equal(t, ok, false)
}
