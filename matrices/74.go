package matrices

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if matrix[0][0] > target {
		return false
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][len(matrix[i])-1] >= target {
			for _, elem := range matrix[i] {
				if elem == target {
					return true
				}
			}
			return false
		}
	}
	return false
}

func Test74() {
	a := searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3)
	fmt.Println(a)

	a = searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13)
	fmt.Println(a)
}
