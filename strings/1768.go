package strings

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	a := make([]byte, 0)
	for i, letter := range word1 {
		a = append(a, byte(letter))
		if i < len(word2) {
			a = append(a, word2[i])
		} else {
			a = append(a, []byte(word1[i+1:])...)
			break
		}
	}
	if len(word2) > len(word1) {
		a = append(a, []byte(word2[len(word1):])...)

	}
	return string(a)
}

func Test1768() {
	a := mergeAlternately("abc", "pqr")
	fmt.Println(a)
	a = mergeAlternately("ab", "pqrs")
	fmt.Println(a)
	a = mergeAlternately("abcd", "pq")
	fmt.Println(a)
}
