package solution

const maxInt = int32(^uint32(0) >> 1)
const minInt = -maxInt - 1

func reverse(x int) int {
	var res int64

	for x != 0 {
		res = res*10 + int64(x)%10
		x /= 10
	}

	if res > int64(maxInt) || res < int64(minInt) {
		return 0
	}

	return int(res)
}
