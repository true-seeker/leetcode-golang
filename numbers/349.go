package numbers

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	type void struct{}
	var m void

	inter := make(map[int]void)
	set1 := make(map[int]void)
	for _, elem := range nums1 {
		set1[elem] = m
	}
	for _, elem := range nums2 {
		if _, ok := set1[elem]; ok {
			inter[elem] = m
		}
	}
	res := make([]int, 0)
	for key := range inter {
		res = append(res, key)
	}
	return res
}

func Test349() {
	a := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Println(a)
	a = intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	fmt.Println(a)
}
