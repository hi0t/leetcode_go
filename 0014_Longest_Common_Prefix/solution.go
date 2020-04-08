package solution

func longestCommonPrefix(strs []string) string {
	res := ""
	j := 0

	if len(strs) == 0 {
		return res
	}

	for {
		var c byte
		for i := 0; i < len(strs); i++ {
			if j >= len(strs[i]) {
				return res
			}
			if c == 0 {
				c = strs[i][j]
				continue
			} else if strs[i][j] != c {
				return res
			}
		}
		res += string(c)
		j++
	}
}
