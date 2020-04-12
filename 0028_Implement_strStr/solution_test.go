package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 0, strStr("a", "a"))
	assert.Equal(t, 2, strStr("hello", "ll"))
}
