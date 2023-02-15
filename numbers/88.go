package numbers

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	firstArrayIndex := 0
	for _, elem := range nums2 {
		for nums1[firstArrayIndex] <= elem && firstArrayIndex < m {
			firstArrayIndex++
		}
		copy(nums1[firstArrayIndex+1:], nums1[firstArrayIndex:])
		nums1[firstArrayIndex] = elem
		m++
	}
}

func Test88() {
	a := []int{1, 2, 3, 0, 0, 0}
	merge(a, 3, []int{2, 5, 6}, 3)
	fmt.Println(a)

	a = []int{1}
	merge(a, 1, []int{}, 0)
	fmt.Println(a)

	a = []int{0}
	merge(a, 0, []int{1}, 1)
	fmt.Println(a)
}
