package isogram

import (
	"unicode"
)

const test = 1

// IsIsogram returns a bool depending on if the input string is an isogram.
func IsIsogram(in string) bool {

	wordMap := make(map[rune]int)
	for _, value := range in {
		if unicode.IsLetter(value) {
			wordMap[unicode.ToLower(value)]++
		}
	}

	output := true
	for _, v := range wordMap {
		if v > 1 {
			output = false
		}
	}

	return output
}