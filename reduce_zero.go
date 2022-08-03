package roman_number

func numberOfSteps(num int) int {
	return foo(num, 0)
}

func foo(num, result int) int {
	if num == 0 {
		return result
	}
	if num&1 == 1 {
		return foo(num-1, result+1)

	} else {
		return foo(num/2, result+1)
	}
}
