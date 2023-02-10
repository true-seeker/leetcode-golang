package strings

import "fmt"

func romanToInt(s string) int {
	alphabet := map[uint8]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	sum := 0
	for i := 0; i < len(s)-1; i++ {
		if alphabet[s[i]] < alphabet[s[i+1]] {
			sum -= alphabet[s[i]]
		} else {
			sum += alphabet[s[i]]
		}
	}
	sum += alphabet[s[len(s)-1]]
	return sum
}

func Test13() {
	a := romanToInt("III")
	fmt.Println(a)
	a = romanToInt("LVIII")
	fmt.Println(a)
	a = romanToInt("MCMXCIV")
	fmt.Println(a)
}
