package strings

import "fmt"

func isIsomorphic(s string, t string) bool {
	fmt.Println()
	m := make(map[uint8]uint8)
	m2 := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 0 && m2[t[i]] == 0 {
			m[s[i]] = t[i]
			m2[t[i]] = s[i]
		} else {
			if m[s[i]] != t[i] || m2[t[i]] != s[i] {
				return false
			}
		}
	}
	return true
}

func Test205() {
	a := isIsomorphic("egg", "add")
	fmt.Println(a)
	a = isIsomorphic("foo", "bar")
	fmt.Println(a)
	a = isIsomorphic("paper", "title")
	fmt.Println(a)
	a = isIsomorphic("bbbaaaba", "aaabbbba")
	fmt.Println(a)
	a = isIsomorphic("badc", "baba")
	fmt.Println(a)
}
