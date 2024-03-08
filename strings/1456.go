package strings

import "fmt"

func maxVowels(s string, k int) int {
	vowelsInWindow := 0

	for i := 0; i < k; i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			vowelsInWindow++
		}
	}
	maxVowelsCount := vowelsInWindow

	for i := k; i < len(s); i++ {

		switch s[i-k] {
		case 'a', 'e', 'i', 'o', 'u':
			vowelsInWindow--
		}

		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			vowelsInWindow++
			if vowelsInWindow > maxVowelsCount {
				maxVowelsCount = vowelsInWindow
				if maxVowelsCount == k {
					return maxVowelsCount
				}
			}
		}
	}
	return maxVowelsCount
}

func Test1456() {
	a := maxVowels("abciiidef", 3)
	fmt.Println(a)
	a = maxVowels("aeiou", 2)
	fmt.Println(a)
	a = maxVowels("leetcode", 3)
	fmt.Println(a)
}
