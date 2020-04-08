package solution

func isValid(s string) bool {
	pair := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var stack []rune

	for _, c := range s {
		if _, ok := pair[c]; ok {
			stack = append(stack, c)
			continue
		}
		n := len(stack) - 1
		if n < 0 {
			return false
		}
		if pair[stack[n]] != c {
			return false
		}
		stack = stack[:n]
	}

	return len(stack) == 0
}
