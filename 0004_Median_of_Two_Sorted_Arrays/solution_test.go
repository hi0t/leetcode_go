package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0002_Add_Two_Numbers(t *testing.T) {
	assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
