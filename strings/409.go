package strings

import (
	"fmt"
)

func longestPalindrome(s string) int {
	lettersCount := map[int32]int{}
	for _, elem := range s {
		lettersCount[elem]++
	}

	answer := 0
	oddCount := 0
	for _, count := range lettersCount {
		if count%2 == 0 {
			answer += count
		} else {
			answer += count - 1
			oddCount++
		}
	}
	if oddCount > 0 {
		answer++
	}

	return answer
}

func Test409() {
	a := longestPalindrome("abccccdd")
	fmt.Println(a)
	a = longestPalindrome("a")
	fmt.Println(a)
	a = longestPalindrome("abb")
	fmt.Println(a)
	a = longestPalindrome("ccc")
	fmt.Println(a)
	a = longestPalindrome("banana")
	fmt.Println(a)
	a = longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth")
	fmt.Println(a)
}
