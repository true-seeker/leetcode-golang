package matrices

import "fmt"

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, customer := range accounts {
		localSum := 0
		for _, account := range customer {
			localSum += account
		}
		if localSum > maxWealth {
			maxWealth = localSum
		}
	}
	return maxWealth
}

func Test1627() {
	a := maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}})
	fmt.Println(a)

	a = maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}})
	fmt.Println(a)

	a = maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}})
	fmt.Println(a)
}
