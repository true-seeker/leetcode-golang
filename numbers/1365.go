package numbers

func smallerNumbersThanCurrent(nums []int) []int {
	var result []int
	for i, e := range nums {
		localCount := 0
		for j, e2 := range nums {
			if i == j {
				continue
			}
			if e2 < e {
				localCount++
			}
		}
		result = append(result, localCount)
	}
	return result
}
