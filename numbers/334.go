package numbers

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}
	minm := math.MaxInt
	mid := math.MaxInt
	for _, elem := range nums {
		if elem > mid {
			return true
		}
		if elem < mid && elem > minm {
			mid = elem
		} else if elem < minm {
			minm = elem
		}
	}
	return false
}

func Test334() {
	a := increasingTriplet([]int{1, 2, 3, 4, 5})
	fmt.Println(a)

	a = increasingTriplet([]int{5, 4, 3, 2, 1})
	fmt.Println(a)

	a = increasingTriplet([]int{2, 1, 5, 0, 4, 6})
	fmt.Println(a)

	a = increasingTriplet([]int{1, 5, 0, 4, 1, 3})
	fmt.Println(a)
}
