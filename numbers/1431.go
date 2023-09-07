package numbers

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
