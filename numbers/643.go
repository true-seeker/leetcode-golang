package numbers

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	maxS := 0
	first := nums[0]
	currentS := 0
	for i := 0; i < k; i++ {
		maxS += nums[i]
	}
	currentS = maxS
	for i := k; i < len(nums); i++ {
		currentS = currentS - first + nums[i]
		if currentS > maxS {
			maxS = currentS
		}
		first = nums[i-k+1]
	}
	return float64(maxS) / float64(k)
}

func Test643() {
	a := findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
	fmt.Println(a)
	a = findMaxAverage([]int{5}, 1)
	fmt.Println(a)
	a = findMaxAverage([]int{1, 2, 3}, 2)
	fmt.Println(a)
}
