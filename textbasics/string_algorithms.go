package textbasics

import "unicode/utf8"

func LenOfWords(words []string) int {
	count := 0
	for _, word := range words {
		count += utf8.RuneCountInString(word)
	}
	return count
}
