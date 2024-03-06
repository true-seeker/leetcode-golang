package numbers

func countPairs(nums []int, target int) int {
	nMap := make(map[int]int)
	count := 0
	for _, e := range nums {
		for key, value := range nMap {
			if key+e < target {
				count += value
			}
		}
		nMap[e]++
	}
	return count
}
