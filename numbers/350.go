package numbers

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	var answer []int
	for _, elem := range nums1 {
		hash[elem]++
	}

	for _, elem := range nums2 {
		if hash[elem] > 0 {
			answer = append(answer, elem)
			hash[elem]--
		}
	}
	return answer
}

func Test350() {
	a := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Println(a)

	a = intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	fmt.Println(a)
}
