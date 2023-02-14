package matrices

import "fmt"

func deleteIsland(grid [][]byte, i, j int) [][]byte {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] == '0' {
		return grid
	}

	grid[i][j] = '0'

	grid = deleteIsland(grid, i-1, j)
	grid = deleteIsland(grid, i, j-1)
	grid = deleteIsland(grid, i+1, j)
	grid = deleteIsland(grid, i, j+1)
	return grid
}

func numIslands(grid [][]byte) int {
	count := 0

	for i, row := range grid {
		for j, elem := range row {
			if elem == '1' {
				grid = deleteIsland(grid, i, j)
				count++
			}
		}
	}
	return count
}

func Test200() {
	a := numIslands([][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0}})
	fmt.Println(a)

	a = numIslands([][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1}})
	fmt.Println(a)
}
