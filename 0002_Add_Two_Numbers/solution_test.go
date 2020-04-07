package solution

import (
	"leetcode/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	l1 := util.NewListNode(2, 4, 3)
	l2 := util.NewListNode(5, 6, 4)

	assert.Equal(t, util.NewListNode(7, 0, 8), addTwoNumbers(l1, l2))
}
