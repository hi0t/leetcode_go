package solution

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	idx := make(map[int]int)
	uniq := make(map[[4]int]bool)

	for i := 1; i < len(nums); i++ {
		idx[nums[i-1]] = i - 1
		for j := i + 1; j < len(nums); j++ {
			for n := j + 1; n < len(nums); n++ {
				k, ok := idx[target-(nums[i]+nums[j]+nums[n])]
				if !ok {
					continue
				}
				arr := [4]int{nums[i], nums[j], nums[n], nums[k]}
				sort.Ints(arr[:])
				if uniq[arr] {
					continue
				}
				uniq[arr] = true
				res = append(res, arr[:])
			}
		}
	}

	return res
}
