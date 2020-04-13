package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []string{"()()()", "()(())", "(())()", "(()())", "((()))"}, generateParenthesis(3))
}
