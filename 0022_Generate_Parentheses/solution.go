package solution

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return []string{""}
	}

	for i := 0; i < n; i++ {
		for _, left := range generateParenthesis(i) {
			for _, right := range generateParenthesis(n - 1 - i) {
				res = append(res, "("+left+")"+right)
			}
		}
	}

	return res
}
