package roman_number

import (
	"fmt"
	"math"
)

func isPerfectSquare(num int) bool {
	var newton = func(x float64) float64 {
		for i, z, old := 1, x/2, 0.0; ; i++ {
			if old, z = z, z-(z*z-x)/2/z; math.Abs(old-z) < 1e-5 {
				fmt.Printf("Ran %v iterations\n", i)
				return z
			}
		}
	}
	foo := newton(float64(num))
	return int(foo)*int(foo) == num
}
