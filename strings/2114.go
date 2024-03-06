package strings

import "strings"

func mostWordsFound(sentences []string) int {
	maxWords := 0
	for _, sentence := range sentences {
		wordsCount := len(strings.Split(sentence, " "))
		if wordsCount > maxWords {
			maxWords = wordsCount
		}
	}
	return maxWords
}
