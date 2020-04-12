package solution

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (r + l) / 2
		if target > nums[m] {
			l = m + 1
		} else if target < nums[m] {
			r = m - 1
		} else {
			return m
		}
	}

	return l
}
