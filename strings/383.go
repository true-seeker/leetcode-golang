package strings

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	lettersCount := map[int32]int{}
	for _, elem := range ransomNote {
		lettersCount[elem]++
	}

	for _, elem := range magazine {
		if lettersCount[elem] != 0 {
			lettersCount[elem]--
		}
	}
	for _, count := range lettersCount {
		if count != 0 {
			return false
		}
	}
	return true
}

func Test383() {
	a := canConstruct("a", "b")
	fmt.Println(a)
	a = canConstruct("aa", "ab")
	fmt.Println(a)
	a = canConstruct("paper", "title")
	fmt.Println(a)
	a = canConstruct("bbbaaaba", "aaabbbba")
	fmt.Println(a)
	a = canConstruct("aa", "aab")
	fmt.Println(a)
}
