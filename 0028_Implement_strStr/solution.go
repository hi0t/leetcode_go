package solution

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	i := 0
	j := 0
	for i+j < len(haystack) {
		if haystack[i+j] == needle[j] {
			if j == len(needle)-1 {
				return i
			}
			j++
			continue
		}
		i++
		j = 0
	}

	return -1
}
