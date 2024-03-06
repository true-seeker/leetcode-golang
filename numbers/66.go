package numbers

import "fmt"

func plusOne(digits []int) []int {
	lastIndex := len(digits) - 1
	if digits[lastIndex] == 9 {
		digits[lastIndex] = 0
		if lastIndex > 0 {
			for i := len(digits) - 2; i >= 0; i-- {
				if digits[i] == 9 {
					digits[i] = 0
					if i == 0 {
						digits = append([]int{1}, digits...)
					}
				} else {
					digits[i]++
					break
				}
			}
		} else {
			digits = append([]int{1}, digits...)
		}
	} else {
		digits[lastIndex]++
	}
	return digits
}

func Test66() {
	a := plusOne([]int{1, 2, 3})
	fmt.Println(a)
	a = plusOne([]int{4, 3, 2, 1})
	fmt.Println(a)
	a = plusOne([]int{9})
	fmt.Println(a)
	a = plusOne([]int{1, 3, 5, 6})
	fmt.Println(a)
	a = plusOne([]int{1, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	fmt.Println(a)
	a = plusOne([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	fmt.Println(a)
}
