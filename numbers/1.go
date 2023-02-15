package numbers

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func Test1() {
	a := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(a)

	a = twoSum([]int{3, 2, 4}, 6)
	fmt.Println(a)

	a = twoSum([]int{3, 3}, 6)
	fmt.Println(a)
}
