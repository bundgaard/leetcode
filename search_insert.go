package roman_number

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func searchInsert(nums []int, target int) int {
	result := 0
	if len(nums) == 1 && target > nums[0] {
		return 1
	}

	for i := 0; i < len(nums)-1; i++ {

		if target == nums[i] {
			result = i
			break
		} else if target > nums[i] && target <= nums[i+1] {
			result = i + 1
			break
		} else if target > nums[i+1] && len(nums)-1 == i+1 {
			result = len(nums)
			break
		} else if target < nums[i] && i == 0 {
			result = 0
			break
		} else if target > nums[i] && target > nums[i+1] {
			continue
		}

	}
	return result

}
