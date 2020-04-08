package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
