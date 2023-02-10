package numbers

import "fmt"

func removeElement(nums []int, val int) int {
	k := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			k--
		}
	}
	fmt.Println(nums)
	return k
}

func Test27() {
	a := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(a)

	a = removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	fmt.Println(a)

	a = removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	fmt.Println(a)
}
