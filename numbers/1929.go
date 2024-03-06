package numbers

import "fmt"

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func Test1929() {
	a := getConcatenation([]int{2, 2, 1})
	fmt.Println(a)
	a = getConcatenation([]int{4, 1, 2, 1, 3})
	fmt.Println(a)
	a = getConcatenation([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(a)
	a = getConcatenation([]int{1, 1})
	fmt.Println(a)
}
