package strings

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	result := make([]string, 0)
	word := make([]byte, 0)
	is_word := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if is_word {
				is_word = false
				temp := make([]byte, len(word))
				for j := len(word) - 1; j >= 0; j-- {
					temp = append(temp, word[j])
				}
				result = append(result, string(temp))
				word = make([]byte, 0)
			}
		} else {
			is_word = true
			word = append(word, s[i])
		}
	}

	if len(word) > 0 {
		temp := make([]byte, len(word))
		for j := len(word) - 1; j >= 0; j-- {
			temp = append(temp, word[j])
		}
		result = append(result, string(temp))
	}
	a := strings.Join(result, string(' '))

	return strings.ReplaceAll(a, "\u0000", "")
}

func Test151() {
	a := reverseWords("the sky is blue")
	fmt.Println(a)
	a = reverseWords("  hello world  ")
	fmt.Println(a)
	a = reverseWords("a good   example")
	fmt.Println(a)
}
