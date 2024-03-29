package solution

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := ""
	cycleLen := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += cycleLen {
			res += string(s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < len(s) {
				res += string(s[j+cycleLen-i])
			}
		}
	}
	return res
}
