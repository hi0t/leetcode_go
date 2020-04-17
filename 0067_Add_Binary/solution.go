package solution

func addBinary(a string, b string) string {
	res := ""
	var carry byte
	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 {
		var sum byte
		if i >= 0 {
			sum += a[i] - '0'
			i--
		}
		if j >= 0 {
			sum += b[j] - '0'
			j--
		}
		sum += carry
		if sum > 1 {
			sum -= 2
			carry = 1
		} else {
			carry = 0
		}
		res = string('0'+sum) + res
	}

	if carry > 0 {
		res = string('0'+carry) + res
	}

	return res
}
