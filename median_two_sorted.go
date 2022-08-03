package roman_number

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, len(nums1)+len(nums2))
	copy(merged, nums1)
	copy(merged[len(nums1):], nums2)

	sort.Ints(merged)

	if len(merged)%2 == 0 {
		var mid = len(merged) / 2
		half1 := merged[mid]
		half2 := merged[mid-1]
		ans := float64(half1+half2) / 2.0
		return ans
	}
	ans := float64(merged[len(merged)/2])
	return ans
}
