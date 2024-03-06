package strings

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	a, b := len(str1), len(str2)
	for b != 0 {
		a, b = b, a%b
	}
	return str1[:a]
}

func Test1071() {
	a := gcdOfStrings("ABCABC", "ABC")
	fmt.Println(a)
	a = gcdOfStrings("ABABAB", "ABAB")
	fmt.Println(a)
	a = gcdOfStrings("LEET", "CODE")
	fmt.Println(a)
}
