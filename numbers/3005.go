package numbers

import "fmt"

func maxFrequencyElements(nums []int) int {
	m := make(map[int]int)
	maxFreq := 0
	res := 0
	for _, elem := range nums {
		m[elem]++
		e := m[elem]
		if e > maxFreq {
			maxFreq = e
			res = e
		} else if e == maxFreq {
			res += e
		}
	}
	return res
}

func Test3005() {
	a := maxFrequencyElements([]int{1, 2, 2, 3, 1, 4})
	fmt.Println(a)
	a = maxFrequencyElements([]int{1, 2, 3, 4, 5})
	fmt.Println(a)
	a = maxFrequencyElements([]int{1, 2, 3})
	fmt.Println(a)
}
