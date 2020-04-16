package solution

import (
	"leetcode/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t,
		util.NewListNode(1, 1, 2, 3, 4, 4, 5, 6),
		mergeKLists([]*util.ListNode{util.NewListNode(1, 4, 5), util.NewListNode(1, 3, 4), util.NewListNode(2, 6)}),
	)
}
