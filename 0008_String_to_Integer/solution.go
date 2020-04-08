package solution

import "math"

type State int

const (
	Start State = iota
	Space
	Sign
	Digit
)

func myAtoi(str string) int {
	res := int64(0)
	s := Start
	sign := int64(1)

loop:
	for _, ch := range str {
		switch {
		case ch == ' ':
			if s == Sign {
				return 0
			} else if s == Digit {
				break loop
			}
			s = Space
		case ch == '-' || ch == '+':
			if s == Sign {
				return 0
			} else if s == Digit {
				break loop
			}
			s = Sign
			if ch == '-' {
				sign *= -1
			}
		case '0' <= ch && ch <= '9':
			res = res*10 + int64(ch-'0')
			if res > math.MaxInt32 {
				res = math.MaxInt32
				if sign == -1 {
					res++
				}
				break loop
			}
			s = Digit
		default:
			if s == Digit {
				break loop
			} else {
				return 0
			}
		}
	}

	res *= sign

	return int(res)
}
