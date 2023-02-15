package numbers

import "fmt"

func climbStairs(n int) int {
	a, b := 1, 1
	for i := n; i > 1; i-- {
		a, b = b, a+b
	}
	return b
}

func Test70() {
	a := climbStairs(2)
	fmt.Println(a)

	a = climbStairs(3)
	fmt.Println(a)

	a = climbStairs(4)
	fmt.Println(a)

	a = climbStairs(15)
	fmt.Println(a)

	a = climbStairs(30)
	fmt.Println(a)
}
