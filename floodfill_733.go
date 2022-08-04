package roman_number

var (
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	fill(image, sr, sc, image[sr][sc], color)
	return image
}

func fill(image [][]int, sr int, sc int, color int, newColor int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != color {
		return
	}
	image[sr][sc] = newColor
	for i := 0; i < 4; i++ {
		fill(image, sr+dx[i], sc+dy[i], color, newColor)
	}
}
