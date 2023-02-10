package strings

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	sIndex := 0
	for i := 0; i < len(t); i++ {
		if s[sIndex] == t[i] {
			sIndex++
			if sIndex == len(s) {
				return true
			}
		}
	}
	return false
}

func Test392() {
	a := isSubsequence("abc", "ahbgdc")
	fmt.Println(a)
	a = isSubsequence("axc", "ahbgdc")
	fmt.Println(a)
	a = isSubsequence("acb", "ahbgdc")
	fmt.Println(a)
}
