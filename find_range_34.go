package roman_number

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	first := firstOccurence(nums, target)
	last := lastOccurence(nums, target)
	return []int{first, last}
}
func lastOccurence(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	result := -1
	for start <= end {
		mid := (start + end) / 2
		n := nums[mid]

		if n <= target {
			start = mid + 1
		} else if n > target {
			end = mid - 1
		}
		if n == target {
			result = mid

		}
	}
	return result
}

func firstOccurence(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	result := -1
	for start <= end {
		mid := (start + end) / 2
		n := nums[mid]
		if n >= target {
			end = mid - 1
		} else if n < target {
			start = mid + 1
		}

		if n == target {
			result = mid
		}
	}
	return result
}
