package numbers

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	count := 0
	for _, e := range hours {
		if e >= target {
			count++
		}
	}
	return count
}
