package roman_number

import "sort"

func kWeakestRows(mat [][]int, k int) []int {

	priority := make([]int, len(mat))

	for i := 0; i < len(mat); i++ {
		soldiers := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				soldiers++
			} else {
				break
			}
		}
		priority[i] = soldiers*1000 + i
	}
	sort.Ints(priority)
	arr := make([]int, len(mat))
	for idx := range priority {
		arr[idx] = priority[idx] % 1000
	}

	return arr[:k]
}
