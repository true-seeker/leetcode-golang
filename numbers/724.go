package numbers

import "fmt"

func pivotIndex(nums []int) int {
	leftSideSum := 0
	for i, elem := range nums {
		localPivotSum := 0
		for j := i + 1; j < len(nums); j++ {
			localPivotSum += nums[j]
		}
		if localPivotSum == leftSideSum {
			return i
		}
		leftSideSum += elem
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
}
