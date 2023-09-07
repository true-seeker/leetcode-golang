package strings

func numJewelsInStones(jewels string, stones string) int {
	jMap := make(map[rune]bool)
	jCount := 0

	for _, j := range jewels {
		jMap[j] = true
	}
	for _, s := range stones {
		if jMap[s] {
			jCount++
		}
	}
	return jCount
}
