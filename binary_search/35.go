package binary_search

import "fmt"

func searchInsert(nums []int, target int) int {
	return binarySearch35(nums, target, 0, len(nums)-1)
}

func binarySearch35(nums []int, target int, left int, right int) int {
	if left > right {
		return left
	}
	middle := (left + right) / 2

	if nums[middle] == target {
		return middle
	}
	if nums[middle] > target {
		return binarySearch35(nums, target, left, middle-1)
	} else {
		return binarySearch35(nums, target, middle+1, right)
	}
}

func Test35() {
	a := searchInsert([]int{1, 3, 5, 6}, 5)
	fmt.Println(a)
	a = searchInsert([]int{1, 3, 5, 6}, 2)
	fmt.Println(a)
	a = searchInsert([]int{1, 3, 5, 6}, 7)
	fmt.Println(a)

}
