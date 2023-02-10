package numbers

import "fmt"

func fizzBuzz(n int) []string {
	var answer []string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer = append(answer, "FizzBuzz")
		} else if i%3 == 0 {
			answer = append(answer, "Fizz")
		} else if i%5 == 0 {
			answer = append(answer, "Buzz")
		} else {
			answer = append(answer, fmt.Sprintf("%d", i))
		}
	}

	return answer
}

func Test412() {
	a := fizzBuzz(3)
	fmt.Println(a)
	a = fizzBuzz(5)
	fmt.Println(a)
	a = fizzBuzz(15)
	fmt.Println(a)
}
