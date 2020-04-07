package solution

func longestPalindrome(s string) string {
	start, end := 0, 0

	if len(s) < 1 {
		return ""
	}

	for i := 0; i < len(s); i++ {
		len1 := mirrored(s, i, i)
		len2 := mirrored(s, i, i+1)
		len := max(len1, len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}

	return s[start : end+1]
}

func mirrored(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		r++
		l--
	}
	return r - l - 1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
