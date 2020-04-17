package solution

func plusOne(digits []int) []int {
	carry := 0

	d := &digits[len(digits)-1]
	*d += 1
	if *d > 9 {
		*d -= 10
		carry = 1
	}

	for i := len(digits) - 2; i >= 0; i-- {
		digits[i] += carry
		if digits[i] > 9 {
			digits[i] -= 10
			carry = 1
		} else {
			carry = 0
		}
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}
