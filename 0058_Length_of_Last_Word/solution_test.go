package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 1, lengthOfLastWord("a "))
	assert.Equal(t, 5, lengthOfLastWord("Hello World"))
	assert.Equal(t, 1, lengthOfLastWord("a"))

}
