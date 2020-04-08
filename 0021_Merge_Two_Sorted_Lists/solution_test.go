package solution

import (
	"leetcode/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, util.NewListNode(1, 1, 2, 3, 4, 4), mergeTwoLists(util.NewListNode(1, 2, 4), util.NewListNode(1, 3, 4)))
}
