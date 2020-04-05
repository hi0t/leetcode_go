package solution

func lengthOfLongestSubstring(s string) int {
	max := 0
	var idx [128]int

	for j, i := 0, 0; j < len(s); j++ {
		if idx[s[j]] > i {
			i = idx[s[j]]
		}
		len := j - i + 1
		if len > max {
			max = len
		}
		idx[s[j]] = j + 1
	}

	return max
}
