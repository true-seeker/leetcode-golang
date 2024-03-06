package numbers

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, e := range nums {
		res ^= e
	}
	return res
}

func Test136() {
	a := singleNumber([]int{2, 2, 1})
	fmt.Println(a)
	a = singleNumber([]int{4, 1, 2, 1, 2})
	fmt.Println(a)
	a = singleNumber([]int{1})
	fmt.Println(a)
}
