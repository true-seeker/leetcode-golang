package numbers

import "fmt"

func reverseVowels(s string) string {
	indicies := make([]int, 0)
	for i, elem := range s {
		switch elem {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			indicies = append(indicies, i)
		}
	}
	s2 := []rune(s)
	for i := 0; i < len(indicies)/2; i++ {
		i1 := indicies[i]
		i2 := indicies[len(indicies)-i-1]
		s2[i1], s2[i2] = s2[i2], s2[i1]
	}
	return string(s2)
}
func Test345() {
	a := reverseVowels("hello")
	fmt.Println(a)

	a = reverseVowels("leetcode")
	fmt.Println(a)
}
