package numbers

import "fmt"

func runningSum(nums []int) []int {
	var running []int
	currentSum := 0
	for _, elem := range nums {
		currentSum += elem
		running = append(running, currentSum)
	}
	return running
}

func Test1480() {
	a := runningSum([]int{1, 2, 3, 4})
	fmt.Println(a)
	a = runningSum([]int{1, 1, 1, 1, 1})
	fmt.Println(a)
	a = runningSum([]int{3, 1, 2, 10, 1})
	fmt.Println(a)
}
