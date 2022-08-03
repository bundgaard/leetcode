package roman_number

import (
	"math"
)

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
var pick = 6

func guess(n int) int {
	if n == pick {
		return 0
	} else if n < pick {
		return 1
	} else {
		return -1
	}
}

func guessNumber(n int) int {
	low := 1
	high := int(math.Pow(2, 31) - 1)
	var mid int
	for low <= high {
		mid = (low + high) / 2
		x := guess(mid)
		if x == 0 {
			return mid
		} else if x == -1 {
			high = mid - 1
		} else if x == 1 {
			low = mid + 1
		}
	}

	return mid
}
