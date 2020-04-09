package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, false, isMatch("ab", ".*c"))
	assert.Equal(t, true, isMatch("aaa", "ab*a*c*a"))
	assert.Equal(t, true, isMatch("aaa", "a*a"))
	assert.Equal(t, false, isMatch("aa", "a"))
	assert.Equal(t, true, isMatch("aa", "a*"))
	assert.Equal(t, true, isMatch("ab", ".*"))
	assert.Equal(t, true, isMatch("aab", "c*a*b"))
	assert.Equal(t, false, isMatch("mississippi", "mis*is*p*."))
	assert.Equal(t, true, isMatch("mississippi", "mis*is*ip*."))
	assert.Equal(t, false, isMatch("aaa", "aaaa"))
}
