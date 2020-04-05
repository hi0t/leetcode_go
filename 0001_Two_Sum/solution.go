package solution

func twoSum(nums []int, target int) []int {
	idx := make(map[int]int)

	for i, n := range nums {
		i0, ok := idx[target-n]
		if ok {
			return []int{i0, i}
		}
		idx[n] = i
	}

	return []int{0, 0}
}
