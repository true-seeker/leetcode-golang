package strings

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	lastLength := 0
	for _, c := range s {
		if c != ' ' {
			length++
		} else {
			if length != 0 {
				lastLength = length
			}
			length = 0
		}
	}
	if length != 0 {
		lastLength = length
	}
	return lastLength
}

func Test58() {
	a := lengthOfLastWord("Hello World")
	fmt.Println(a)
	a = lengthOfLastWord("   fly me   to   the moon  ")
	fmt.Println(a)
	a = lengthOfLastWord("luffy is still joyboy")
	fmt.Println(a)
	a = lengthOfLastWord("amogus")
	fmt.Println(a)
	a = lengthOfLastWord("badc sus ")
	fmt.Println(a)
}
