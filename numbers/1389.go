package numbers

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	var target []int
	for i, e := range index {
		target = append(target, 0)
		copy(target[e+1:], target[e:])
		target[e] = nums[i]
	}
	return target
}

func Test1389() {
	a := createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1})
	fmt.Println(a)

	//a = createTargetArray([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	//fmt.Println(a)
}
