package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "11110", addBinary("1111", "1111"))
	assert.Equal(t, "10101", addBinary("1010", "1011"))
}
