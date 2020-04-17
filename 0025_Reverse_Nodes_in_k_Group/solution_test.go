package solution

import (
	"leetcode/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, util.NewListNode(2, 1, 4, 3, 5), reverseKGroup(util.NewListNode(1, 2, 3, 4, 5), 2))
	assert.Equal(t, util.NewListNode(3, 2, 1, 4, 5), reverseKGroup(util.NewListNode(1, 2, 3, 4, 5), 3))
}
