package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "III", intToRoman(3))
	assert.Equal(t, "MCMXCIV", intToRoman(1994))
}
