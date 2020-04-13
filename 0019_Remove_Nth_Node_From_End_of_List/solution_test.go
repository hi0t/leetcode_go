package solution

import (
	"leetcode/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, nil, removeNthFromEnd(util.NewListNode(1), 1))
	assert.Equal(t, util.NewListNode(1, 2, 3, 5), removeNthFromEnd(util.NewListNode(1, 2, 3, 4, 5), 2))
}
