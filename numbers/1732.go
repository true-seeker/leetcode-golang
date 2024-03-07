package numbers

import "fmt"

func largestAltitude(gain []int) int {
	maxAlt := 0
	currentAlt := 0
	for _, elem := range gain {
		currentAlt += elem
		if currentAlt > maxAlt {
			maxAlt = currentAlt
		}
	}
	return maxAlt
}

func Test1732() {
	a := largestAltitude([]int{-5, 1, 5, 0, -7})
	fmt.Println(a)
	a = largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2})
	fmt.Println(a)
}
