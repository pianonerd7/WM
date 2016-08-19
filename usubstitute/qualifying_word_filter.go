package usubstitute

import (
	"strings"
)

func didPassAllFilterForSynset(newSense string, uniqueSenses []string) bool {
	if !isStringInSlice(newSense, uniqueSenses) && isOneWord(newSense) {
		return false
	}
	return true
}

func isQualifyingWord(word string) bool {
	return isOneWord(word)
}

func isStringInSlice(word string, list []string) bool {
	for _, sliceElement := range list {
		if word == sliceElement {
			return true
		}
	}
	return false
}

func isOneWord(word string) bool {
	if strings.ContainsAny(word, "_ ( ) 1 2 3 4 5 6 7 8 9 0 ' ") {
		return false
	}
	return true
}
