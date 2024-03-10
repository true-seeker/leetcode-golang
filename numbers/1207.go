package numbers

import "fmt"

func uniqueOccurrences(arr []int) bool {
	type void struct{}
	var m void
	count := make(map[int]int)
	uniqueCount := make(map[int]void)
	for _, elem := range arr {
		count[elem]++
	}
	for _, value := range count {
		if _, ok := uniqueCount[value]; ok {
			return false
		}
		uniqueCount[value] = m
	}
	return true
}

func Test1207() {
	a := uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})
	fmt.Println(a)
	a = uniqueOccurrences([]int{1, 2})
	fmt.Println(a)
	a = uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})
	fmt.Println(a)
}
