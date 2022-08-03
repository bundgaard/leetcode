package roman_number

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 0; i < n; i++ {
		f := i + 1
		if f%5 == 0 && f%3 == 0 {
			result[i] = "FizzBuzz"
		} else if f%5 == 0 {
			result[i] = "Buzz"
		} else if f%3 == 0 {
			result[i] = "Fizz"
		} else {
			result[i] = strconv.Itoa(f)
		}
	}

	return result
}
