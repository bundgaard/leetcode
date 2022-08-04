package roman_number

func nextGreatestLetter(letters []byte, target byte) byte {

	n := len(letters)
	if target >= letters[n-1] {
		return letters[0]
	}

	result := 0
	l := 0
	r := n - 1
	for l <= r {
		mid := (l + r) / 2
		if letters[mid] > target {
			r = mid - 1
			result = mid
		} else {
			l = mid + 1
		}
	}
	return letters[result]
}
