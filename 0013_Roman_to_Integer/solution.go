package solution

func romanToInt(s string) int {
	res := 0

	t := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	for i := 0; i < len(s); i++ {
		var c string
		var d int

		if i+1 < len(s) {
			c = s[i : i+2]
			d = t[c]
		}

		if d == 0 {
			d = t[string(s[i])]
		} else {
			i++
		}

		res += d
	}

	return res
}
