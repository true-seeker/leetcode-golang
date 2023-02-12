package binary_search

import "fmt"

func binarySearch(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	middle := (left + right) / 2

	if nums[middle] == target {
		return middle
	}
	if nums[middle] > target {
		return binarySearch(nums, target, left, middle-1)
	} else {
		return binarySearch(nums, target, middle+1, right)

	}
}

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func Test704() {
	a := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println(a)
	a = search([]int{-1, 0, 3, 5, 9, 12}, 2)
	fmt.Println(a)
}
