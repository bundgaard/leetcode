package roman_number

func climbStairs(n int) int {
	var arr [46]int
	arr[1] = 1
	arr[2] = 2

	for i := 3; i < 46; i++ {
		arr[i] = arr[i-2] + arr[i-1]
	}
	return arr[n]
}
