package numbers

func numIdenticalPairs(nums []int) int {
	nMap := make(map[int]int)
	pairsCount := 0

	for _, e := range nums {
		pairsCount += nMap[e]
		nMap[e]++
	}
	return pairsCount
}
