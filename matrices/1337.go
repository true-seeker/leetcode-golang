package matrices

import (
	"fmt"
)

func kWeakestRows(mat [][]int, k int) []int {
	var power []int
	var powerIndicies []int
	for i, elem := range mat {
		localPower := 0
		for j := 0; j < len(elem) && elem[j] == 1; j++ {
			localPower++
		}

		if len(power) == 0 {
			power = append(power, localPower)
			powerIndicies = append(powerIndicies, i)
		} else {
			j := 0
			for j = 0; j < len(power); j++ {
				if power[j] > localPower {
					power = append(power, 0)
					copy(power[j+1:], power[j:])
					power[j] = localPower

					powerIndicies = append(powerIndicies, 0)
					copy(powerIndicies[j+1:], powerIndicies[j:])
					powerIndicies[j] = i
					break
				}
			}
			if j == len(power) {
				power = append(power, localPower)
				powerIndicies = append(powerIndicies, i)
			}
		}
	}
	return powerIndicies[:k]
}

func Test1337() {
	a := kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}},
		3)
	fmt.Println(a)
}
