package numbers

import (
	"fmt"
)

func maxArea(height []int) int {
	maxVol := 0
	start := 0
	end := len(height) - 1
	for start < end {
		leftP := height[start]
		rightP := height[end]
		vol := 0
		if leftP < rightP {
			vol = (end - start) * leftP
			start++
		} else {
			vol = (end - start) * rightP
			end--
		}
		if vol > maxVol {
			maxVol = vol
		}
	}
	return maxVol
}

func Test11() {
	a := maxArea([]int{2, 2, 1})
	fmt.Println(a)
	a = maxArea([]int{4, 1, 2, 1, 3})
	fmt.Println(a)
	a = maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(a)
	a = maxArea([]int{1, 1})
	fmt.Println(a)
}
