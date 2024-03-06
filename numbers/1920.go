package numbers

func buildArray(nums []int) []int {
	var ans []int
	for _, e := range nums {
		ans = append(ans, nums[e])
	}
	return ans
}
