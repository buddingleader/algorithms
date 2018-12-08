package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	a  = []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	b  = []int{16, 15, 10, 8, 14, 9, 3, 2, 4, 1, 7}
	c1 = []int{14, 10, 8, 7, 9, 3, 2, 4, 1}
	c2 = []int{10, 8, 7, 9, 3, 2, 4, 1}
	c3 = []int{9, 7, 8, 3, 2, 4, 1}
	d  = []int{16, 15, 10, 14, 7, 9, 3, 2, 8, 1}
)

func Test_Insert(t *testing.T) {
	assert.Equal(t, b, Insert(a, 15))
}

func Test_ExtractMax(t *testing.T) {
	a, max := ExtractMax(a)
	assert.Equal(t, a, c1)
	assert.Equal(t, max, 16)
	a, max = ExtractMax(a)
	assert.Equal(t, a, c2)
	assert.Equal(t, max, 14)
	a, max = ExtractMax(a)
	assert.Equal(t, a, c3)
	assert.Equal(t, max, 10)
}

func Test_IncreaseKey(t *testing.T) {
	IncreaseKey(a, 8, 15)
	assert.Equal(t, d, a)
}
