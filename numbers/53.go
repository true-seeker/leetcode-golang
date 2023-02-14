package numbers

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	max, sum := -math.MaxInt, -math.MaxInt
	for _, elem := range nums {
		if sum < 0 {
			sum = elem
		} else {
			sum += elem
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func Test53() {
	a := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(a)

	a = maxSubArray([]int{1})
	fmt.Println(a)

	a = maxSubArray([]int{5, 4, -1, 7, 8})
	fmt.Println(a)
}
