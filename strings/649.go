package strings

import (
	"fmt"
	"strings"
)

func predictPartyVictory(senate string) string {
	rc := 0
	dc := 0
	ind := 0
	for _, elem := range senate {
		switch elem {
		case 'R':
			rc++
		case 'D':
			dc++
		}
	}
	j := len(senate)
	for rc > 0 || dc > 0 {
		for i := 0; i < j; i++ {
			switch senate[i] {
			case 'R':
				if dc > 0 {
					ind = strings.Index(senate[i:], "D")
					if ind != -1 {
						senate = senate[:ind+i] + senate[ind+i+1:]
					} else {
						ind = strings.Index(senate, "D")
						senate = senate[:ind] + senate[ind+1:]
					}
					dc--
				} else {
					return "Radiant"
				}
			case 'D':
				if rc > 0 {
					ind = strings.Index(senate[i:], "R")
					if ind != -1 {
						senate = senate[:ind+i] + senate[ind+i+1:]
					} else {
						ind = strings.Index(senate, "R")
						senate = senate[:ind] + senate[ind+1:]
					}
					rc--
				} else {
					return "Dire"
				}
			}
			j = len(senate)

		}
	}
	return "haha"
}
func Test649() {
	a := predictPartyVictory("RD")
	fmt.Println(a)
	a = predictPartyVictory("RDD")
	fmt.Println(a)
	a = predictPartyVictory("DDRRR")
	fmt.Println(a)
	a = predictPartyVictory("DRRD")
	fmt.Println(a)
}
