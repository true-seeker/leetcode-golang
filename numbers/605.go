package numbers

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	skip := 0
	for i := 0; i < len(flowerbed) && n != 0; i++ {
		if skip > 0 {
			skip--
			continue
		}
		if flowerbed[i] == 1 {
			skip = 1
			continue
		}
		if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
			skip = 2
			continue
		}
		n--
		skip = 1
	}
	return n == 0
}

func Test605() {
	a := canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 2)
	fmt.Println(a)
	a = canPlaceFlowers([]int{0, 0, 1, 0, 0}, 1)
	fmt.Println(a)
	a = canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)
	fmt.Println(a)
	a = canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)
	fmt.Println(a)
	a = canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2)
	fmt.Println(a)
	a = canPlaceFlowers([]int{1}, 1)
	fmt.Println(a)
	a = canPlaceFlowers([]int{0}, 1)
	fmt.Println(a)
}
