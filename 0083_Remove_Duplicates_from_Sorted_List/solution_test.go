package solution

import (
	"leetcode/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, util.NewListNode(1, 2, 3), deleteDuplicates(util.NewListNode(1, 1, 1, 2, 3, 3)))
}
