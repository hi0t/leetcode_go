package solution

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	m := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	if len(digits) == 0 {
		return res
	}

	res = append(res, "")
	for _, c := range digits {
		res = expand(res, m[c])
	}

	return res
}

func expand(l []string, a string) []string {
	next := make([]string, 0)
	for _, s := range l {
		for _, c := range a {
			next = append(next, s+string(c))
		}
	}
	return next
}
