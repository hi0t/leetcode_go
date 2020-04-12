package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, [][]int{{-1, 0, 1}, {-1, -1, 2}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
