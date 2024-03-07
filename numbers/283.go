package numbers

func moveZeroes(nums []int) {
	zeroCount := 0
	j := len(nums)
	for i := 0; i < j; {
		if nums[i] == 0 {
			zeroCount++
			nums = append(nums[:i], nums[i+1:]...)
			j = len(nums)
		} else {
			i++
		}
	}
	for i := 0; i < zeroCount; i++ {
		nums = append(nums, 0)
	}
}
func Test283() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0})
	moveZeroes([]int{2, 1, 5, 0, 4, 6})
	moveZeroes([]int{1, 5, 4, 1, 3, 0})
}
