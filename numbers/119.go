package numbers

import "fmt"

func getRow(rowIndex int) []int {
	result := []int{1}
	bound := 0
	if rowIndex == 0 {
		return result
	}
	index := 1
	if rowIndex%2 != 0 {
		bound = rowIndex/2 + 1
	} else {
		bound = rowIndex / 2
	}
	value := 1
	for i := rowIndex; i > bound; i-- {
		value *= i
		value /= index
		index++
		result = append(result, value)
	}
	if rowIndex%2 == 0 {
		bound = len(result) - 2
	} else {
		bound = len(result) - 1
	}
	for i := bound; i >= 0; i-- {
		result = append(result, result[i])
	}
	return result
}

func Test119() {
	a := getRow(0)
	fmt.Println(a)
	a = getRow(1)
	fmt.Println(a)
	a = getRow(2)
	fmt.Println(a)
	a = getRow(3)
	fmt.Println(a)
	a = getRow(4)
	fmt.Println(a)
	a = getRow(5)
	fmt.Println(a)
	a = getRow(6)
	fmt.Println(a)
	a = getRow(7)
	fmt.Println(a)
	a = getRow(21)
	fmt.Println(a)
	a = getRow(30)
	fmt.Println(a)
}
