package roman_number

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	low := mySqrt(x>>2) << 1
	high := low + 1
	if high*high > x {
		return low
	}
	return high
}
