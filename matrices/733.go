package matrices

import "fmt"

func makeFlood(image [][]int, i, j, defColor, color int) [][]int {
	if i < 0 || j < 0 || i >= len(image) || j >= len(image[i]) || image[i][j] != defColor {
		return image
	}

	image[i][j] = color

	image = makeFlood(image, i-1, j, defColor, color)
	image = makeFlood(image, i, j-1, defColor, color)
	image = makeFlood(image, i+1, j, defColor, color)
	image = makeFlood(image, i, j+1, defColor, color)

	return image
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	return makeFlood(image, sr, sc, image[sr][sc], color)
}

func Test733() {
	a := floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2)
	fmt.Println(a)

	a = floodFill([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 1, 1, 0)
	fmt.Println(a)
}
