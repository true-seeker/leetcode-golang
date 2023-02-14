package numbers

import "fmt"

func containsDuplicate(nums []int) bool {
	set := make(map[int]int)
	for _, elem := range nums {
		set[elem]++
		if set[elem] == 2 {
			return true
		}
	}
	return false
}

func Test217() {
	a := containsDuplicate([]int{1, 2, 3, 1})
	fmt.Println(a)

	a = containsDuplicate([]int{1, 2, 3, 4})
	fmt.Println(a)

	a = containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	fmt.Println(a)
}
