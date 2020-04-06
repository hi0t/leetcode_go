package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res float64
	nums := make([]int, len(nums1)+len(nums2))
	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums[i+j] = nums1[i]
			i++
		} else {
			nums[i+j] = nums2[j]
			j++
		}
	}
	for i < len(nums1) {
		nums[i+j] = nums1[i]
		i++
	}
	for j < len(nums2) {
		nums[i+j] = nums2[j]
		j++
	}

	len := len(nums)
	if len%2 == 0 {
		res = float64(nums[len/2-1]+nums[len/2]) / 2.0
	} else {
		res = float64(nums[len/2])
	}

	return res
}
