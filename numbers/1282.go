package numbers

func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][][]int)
	result := make([][]int, 0)
	for i, e := range groupSizes {
		groupList := groups[e]
		if groupList == nil {
			groupList = make([][]int, 1)
		}

		group := groupList[len(groupList)-1]
		if len(group) == e {
			groupList = append(groupList, []int{})
			group = []int{}
		}
		group = append(group, i)
		groupList[len(groupList)-1] = group
		groups[e] = groupList
	}
	for _, value := range groups {
		for _, group := range value {
			result = append(result, group)
		}
	}
	return result
}
