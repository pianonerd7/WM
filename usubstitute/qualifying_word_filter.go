package usubstitute

import (
	"strings"
)

func didPassAllFilter(newSense string, uniqueSenses []string) bool {
	if !isStringInSlice(newSense, uniqueSenses) && isOneWord(newSense) {
		return false
	}
	return true
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
	if strings.Contains(word, "_") || strings.Contains(word, "(") || strings.Contains(word, ")") || strings.Contains(word, "2") {
		return false
	}
	return true
}
