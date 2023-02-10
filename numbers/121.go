package numbers

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minBuy := math.MaxInt
	globalMaxDiff := 0
	for i, elem := range prices {
		if elem >= minBuy {
			continue
		}
		minBuy = elem
		localMax := 0

		for _, elem2 := range prices[i:] {
			if elem2 > localMax {
				localMax = elem2
			}
		}
		diff := localMax - minBuy
		if diff > globalMaxDiff {
			globalMaxDiff = diff
		}
	}
	return globalMaxDiff
}

func Test121() {
	a := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(a)

	a = maxProfit([]int{7, 6, 4, 3, 1})
	fmt.Println(a)

	a = maxProfit([]int{2, 4, 1})
	fmt.Println(a)

	a = maxProfit([]int{6, 1, 3, 2, 4, 7})
	fmt.Println(a)
}
