package numbers

import "fmt"

func pivotIndex(nums []int) int {
	left := 0
	right := 0

	for _, num := range nums {
		right += num
	}

	for i, num := range nums {
		if left == right-num {
			return i
		}
		left += num
		right -= num
	}

	return -1
}

func Test724() {
	a := pivotIndex([]int{1, 7, 3, 6, 5, 6})
	fmt.Println(a)
	a = pivotIndex([]int{1, 2, 3})
	fmt.Println(a)
	a = pivotIndex([]int{2, 1, -1})
	fmt.Println(a)
	a = pivotIndex([]int{-1, -1, -1, -1, -1, 0})
	fmt.Println(a)
}
