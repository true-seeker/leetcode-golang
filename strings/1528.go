package strings

func restoreString(s string, indices []int) string {
	result := []rune(s)
	for i, e := range indices {
		result[e] = rune(s[i])
	}
	return string(result)
}
