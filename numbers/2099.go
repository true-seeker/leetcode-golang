package numbers

import (
	"fmt"
	"sort"
)

func maxSubsequence(nums []int, k int) []int {
	sortedNums := make([]int, len(nums))
	var answer []int
	copy(sortedNums, nums)

	sort.Slice(sortedNums, func(i, j int) bool {
		return sortedNums[i] > sortedNums[j]
	})

	sortedNums = sortedNums[:k]
	for _, elem := range nums {
		for j, ans := range sortedNums {
			if ans == elem {
				answer = append(answer, elem)
				sortedNums = append(sortedNums[:j], sortedNums[j+1:]...)
				break
			}
		}
		if len(answer) == k {
			return answer
		}
	}
	return answer
}

func Test2099() {
	a := maxSubsequence([]int{3, 4, 3, 3}, 2)
	fmt.Println(a)
}
