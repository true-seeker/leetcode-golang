package numbers

import "fmt"

type void struct{}

var m void

func findDifference(nums1 []int, nums2 []int) [][]int {
	res1 := make([]int, 0)
	res2 := make([]int, 0)
	union := make(map[int]void)
	intersection := make(map[int]void)
	set1 := make(map[int]void)
	set2 := make(map[int]void)
	for _, elem := range nums1 {
		union[elem] = m
		set1[elem] = m
	}
	for _, elem := range nums2 {
		union[elem] = m
		set2[elem] = m
	}
	for key := range union {
		_, ok1 := set1[key]
		_, ok2 := set2[key]
		if ok1 && ok2 {
			intersection[key] = m
		}
	}
	for key := range set1 {
		if _, ok := intersection[key]; !ok {
			res1 = append(res1, key)
		}
	}

	for key := range set2 {
		if _, ok := intersection[key]; !ok {
			res2 = append(res2, key)
		}
	}
	return [][]int{res1, res2}
}

func Test2215() {
	a := findDifference([]int{1, 2, 3}, []int{2, 4, 6})
	fmt.Println(a)
	a = findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2})
	fmt.Println(a)
}
