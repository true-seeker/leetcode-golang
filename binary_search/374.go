package binary_search

import "fmt"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {
	pick := 2
	if num > pick {
		return -1
	}
	if num < pick {
		return 1
	}
	return 0
}

func guessNumber(n int) int {
	res := -1
	l, r := 1, n+1
	mid := 0
	for res != 0 {
		mid = (l + r) / 2
		res = guess(mid)
		switch res {
		case -1:
			r = mid
		case 1:
			l = mid
		case 0:
			return mid
		}
	}
	return mid
}

func Test374() {
	a := guessNumber(2)
	fmt.Println(a)
}
