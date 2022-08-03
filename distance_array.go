package roman_number

import (
	"fmt"
	"strconv"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var abs = func(a int) int {
		if a < 0 {
			return -1 * a
		}
		return a
	}
	count := 0
	for i := 0; i < len(arr1); i++ {
		fmt.Println("for arr1[" + strconv.Itoa(i) + "]=" + strconv.Itoa(arr1[i]))
		acc := true
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				acc = false
				break
			}

		}
		if acc {
			count++
		}
	}
	return count
}
