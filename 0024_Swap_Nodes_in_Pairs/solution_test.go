package solution

import (
	"leetcode/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, util.NewListNode(2, 1, 4, 3), swapPairs(util.NewListNode(1, 2, 3, 4)))
}
