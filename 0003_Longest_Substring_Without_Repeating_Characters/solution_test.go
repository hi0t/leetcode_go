package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0002_Add_Two_Numbers(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 1, lengthOfLongestSubstring(" "))
	assert.Equal(t, 2, lengthOfLongestSubstring("au"))
}
