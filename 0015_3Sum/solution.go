package solution

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	idx := make(map[int]int)
	uniq := make(map[[3]int]bool)

	for i := 1; i < len(nums); i++ {
		idx[nums[i-1]] = i - 1
		for j := i + 1; j < len(nums); j++ {
			k, ok := idx[0-(nums[i]+nums[j])]
			if !ok {
				continue
			}
			arr := [3]int{nums[i], nums[j], nums[k]}
			sort.Ints(arr[:])
			if uniq[arr] {
				continue
			}
			uniq[arr] = true
			res = append(res, arr[:])
		}
	}

	return res
}
