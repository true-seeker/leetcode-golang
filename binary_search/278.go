package binary_search

import "fmt"

func isBadVersion(version int) bool {
	if version >= 5 {
		return true
	}
	return false
}

func binarySearch278(left int, right int) int {
	middle := (left + right) / 2

	if left > right {
		return left
	}
	if isBadVersion(middle) {
		return binarySearch278(left, middle-1)
	} else {
		return binarySearch278(middle+1, right)
	}
}

func firstBadVersion(n int) int {
	return binarySearch278(1, n)
}

func Test278() {
	a := firstBadVersion(5)
	fmt.Println(a)
	//a = firstBadVersion(2)
	//fmt.Println(a)
}
