package numbers

func mathMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxIncreaseKeepingSkyline(grid [][]int) int {
	result := 0
	maxHeightsRow := make([]int, len(grid))
	maxHeightsColumn := make([]int, len(grid))

	for i := 0; i < len(grid); i++ {
		rowMax := 0
		for j := 0; j < len(grid); j++ {
			if grid[i][j] > rowMax {
				rowMax = grid[i][j]
			}
		}
		maxHeightsRow[i] = rowMax
	}
	for i := 0; i < len(grid); i++ {
		columnMax := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > columnMax {
				columnMax = grid[j][i]
			}
		}
		maxHeightsColumn[i] = columnMax
	}

	for i := 0; i < len(grid); i++ {
		maxRowHeight := maxHeightsRow[i]
		for j := 0; j < len(grid); j++ {
			maxColumnHeight := maxHeightsColumn[j]
			minHeightToIncrease := mathMin(maxRowHeight, maxColumnHeight)
			result += minHeightToIncrease - grid[i][j]
		}
	}
	return result
}
