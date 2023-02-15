package numbers

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	prev := 0
	cur := 1
	for i := 2; i <= n; i++ {
		temp := cur
		cur += prev
		prev = temp
	}
	return cur
}

func Test509() {
	a := fib(2)
	fmt.Println(a)

	a = fib(3)
	fmt.Println(a)

	a = fib(4)
	fmt.Println(a)

	a = fib(15)
	fmt.Println(a)

	a = fib(30)
	fmt.Println(a)
}
