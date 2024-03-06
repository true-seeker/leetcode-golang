package strings

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, e := range operations {
		if e[1] == '-' {
			x--
		} else {
			x++
		}
	}
	return x
}
