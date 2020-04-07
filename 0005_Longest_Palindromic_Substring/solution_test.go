package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0002_Add_Two_Numbers(t *testing.T) {
	assert.Equal(t, "aba", longestPalindrome("babad"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
}
