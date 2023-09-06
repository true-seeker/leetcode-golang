package numbers

import "fmt"

func maxArea(height []int) int {
	maxVol := 0
	start := 0
	end := len(height) - 1
	for start < end {
		elem1 := height[start]
		elem2 := height[end]
		vol := 0
		if elem1 < elem2 {
			vol = (end - start) * elem1
			start++
		} else {
			vol = (end - start) * elem2
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
