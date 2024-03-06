package strings

func balancedStringSplit(s string) int {
	balance := 0
	count := 0
	for _, e := range s {
		if e == 'R' {
			balance++
		} else {
			balance--
		}
		if balance == 0 {
			count++
		}
	}
	return count
}
