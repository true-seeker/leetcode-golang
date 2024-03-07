package numbers

import "fmt"

func maxOperations(nums []int, k int) int {
	op := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= k {
			continue
		}
		if m[nums[i]] > 0 {
			m[nums[i]]--
			op++
		} else {
			m[k-nums[i]]++
		}
	}
	return op
}

func Test1679() {
	a := maxOperations([]int{1, 2, 3, 4}, 5)
	fmt.Println(a)
	a = maxOperations([]int{1, 2, 4, 3}, 5)
	fmt.Println(a)
	a = maxOperations([]int{3, 1, 3, 4, 3}, 6)
	fmt.Println(a)
	a = maxOperations([]int{3, 2, 3, 4, 3}, 6)
	fmt.Println(a)
}
