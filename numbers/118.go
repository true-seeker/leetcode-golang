package numbers

import "fmt"

func generate(numRows int) [][]int {
	result := [][]int{{1}}
	for i := 0; i < numRows-1; i++ {
		temp := []int{1}
		for j := 0; j <= i-1; j++ {
			temp = append(temp, result[i][j]+result[i][j+1])
		}
		temp = append(temp, 1)
		result = append(result, temp)
	}
	return result
}

func Test118() {
	a := generate(1)
	fmt.Println(a)
	a = generate(2)
	fmt.Println(a)
	a = generate(3)
	fmt.Println(a)
	a = generate(4)
	fmt.Println(a)
	a = generate(5)
	fmt.Println(a)
	a = generate(6)
	fmt.Println(a)
	a = generate(7)
	fmt.Println(a)
	a = generate(20)
	fmt.Println(a)

}
