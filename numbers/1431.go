package numbers

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxValue := 0
	result := make([]bool, len(candies))
	for _, e := range candies {
		if e > maxValue {
			maxValue = e
		}
	}
	for i, e := range candies {
		if e+extraCandies >= maxValue {
			result[i] = true
		}
	}
	return result
}

func Test1431() {
	a := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	fmt.Println(a)
	a = kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)
	fmt.Println(a)
	a = kidsWithCandies([]int{12, 1, 12}, 10)
	fmt.Println(a)
}
