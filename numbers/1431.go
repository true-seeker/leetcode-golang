package numbers

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	m := 0
	a := make([]int, len(candies))
	b := make([]bool, len(candies))

	for i, elem := range candies {
		a[i] = elem + extraCandies
		if elem > m {
			m = elem
		}
	}
	for i := range a {
		b[i] = a[i] >= m
	}
	return b
}

func Test1431() {
	a := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	fmt.Println(a)
	a = kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)
	fmt.Println(a)
	a = kidsWithCandies([]int{12, 1, 12}, 10)
	fmt.Println(a)
}
