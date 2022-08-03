package roman_number

func canConstruct(ransomNote string, magazine string) bool {

	magazineLetters := make(map[rune]int)
	for _, b := range magazine {
		magazineLetters[b]++
	}

	for _, b := range ransomNote {
		v, ok := magazineLetters[b]

		if ok && v > 0 {
			magazineLetters[b]--
		} else {
			return false
		}
	}

	return true
}
