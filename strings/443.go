package strings

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	currentElem := chars[0]
	firstGroupElemIndex := 0
	groupSize := 1
	j := len(chars)
	oldLen := len(chars)
	for i := 0; i < j-1; {
		i++
		if chars[i] == currentElem {
			groupSize++
		} else {
			if groupSize > 1 {
				oldLen = len(chars)
				chars = append(chars[:firstGroupElemIndex+1], append([]byte(strconv.Itoa(groupSize)), chars[i:]...)...)
				j = len(chars)
				i = i - (oldLen - j)
			}
			firstGroupElemIndex = i
			currentElem = chars[i]
			groupSize = 1
		}
	}
	if groupSize > 1 {
		chars = append(chars[:firstGroupElemIndex+1], []byte(strconv.Itoa(groupSize))...)
	}
	//fmt.Println(string(chars))
	return len(chars)
}

func Test443() {
	a := compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'})
	fmt.Println(a)
	a = compress([]byte{'a'})
	fmt.Println(a)
	a = compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'})
	fmt.Println(a)
	a = compress([]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'})
	fmt.Println(a)
	a = compress([]byte{'a', 'a', 'a', 'a', 'b', 'a'})
	fmt.Println(a)
}
