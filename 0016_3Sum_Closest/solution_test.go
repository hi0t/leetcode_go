package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
