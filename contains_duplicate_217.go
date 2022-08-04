package roman_number

import "log"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]int)

	for _, num := range nums {
		seen[num]++
	}
	log.Println(seen)
	for _, v := range seen {
		if v > 1 {
			return true
		}
	}
	return false
}
