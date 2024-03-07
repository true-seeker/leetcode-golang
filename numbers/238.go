package numbers

import "fmt"

func productExceptSelf(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	res[0], res[size-1] = 1, 1
	for i := 1; i < size; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := size - 2; i >= 0; i-- {
		rightProduct *= nums[i+1]
		res[i] *= rightProduct
	}

	return res
}

func Test238() {
	a := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Println(a)
	a = productExceptSelf([]int{-1, 1, 0, -3, 3})
	fmt.Println(a)
}
