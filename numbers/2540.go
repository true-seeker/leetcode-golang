package numbers

import "fmt"

func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums2[j] < nums1[i] {
			j++
		} else {
			return nums1[i]
		}
	}
	return -1
}

func Test2540() {
	a := getCommon([]int{1, 2, 3}, []int{2, 4})
	fmt.Println(a)
	a = getCommon([]int{1, 2, 3, 6}, []int{2, 3, 4, 5})
	fmt.Println(a)
}
