package numbers

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func leftRightDifference(nums []int) []int {
	leftSum := 0
	rightSum := 0
	answer := make([]int, len(nums))
	for _, e := range nums {
		rightSum += e
	}
	for i, e := range nums {
		rightSum -= e
		answer[i] = abs(leftSum - rightSum)
		leftSum += e
	}
	return answer
}
