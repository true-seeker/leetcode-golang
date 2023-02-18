package matrices

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	reshaped := make([][]int, r)
	for i := 0; i < r; i++ {
		reshaped[i] = make([]int, c)
	}
	var x, y int
	for _, row := range mat {
		for _, elem := range row {
			if x == r {
				return mat
			}
			reshaped[x][y] = elem
			y++
			if y == c {
				y = 0
				x++
			}
		}
	}
	if x != r || y != 0 {
		return mat
	}
	return reshaped
}

func Test566() {
	a := matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4)
	fmt.Println(a)
	a = matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4)
	fmt.Println(a)
	a = matrixReshape([][]int{{1, 2}}, 1, 1)
	fmt.Println(a)
}
