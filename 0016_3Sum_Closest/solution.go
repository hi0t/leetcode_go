package solution

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minLen := math.MaxInt32
	res := 0

	for i := 0; i < len(nums); i++ {
		l := i + 1
		r := len(nums) - 1
		curr := target - nums[i]
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			n := abs(target - s)
			if n < minLen {
				minLen = n
				res = s
			}
			if nums[l]+nums[r] > curr {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
