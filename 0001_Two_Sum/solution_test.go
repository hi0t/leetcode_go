package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	assert.Equal(t, []int{0, 1}, twoSum(nums, target))
}
