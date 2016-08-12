package usubstitute

import (
	"regexp"
	"strings"

	"syntacticsub/wordnet"
)

// GetSynset takes in a word and returns a synset
func GetSynset(word string) []string {
	resultString := wordnet.FindTheInfo_ds(word, 1, 5, 0)

	delimeterRules := func(t rune) bool {
		return t == '{' || t == '}' || t == ','
	}

	return strings.FieldsFunc(resultString, delimeterRules)
}

// MessageToWords takes in a string of words representing a message
// and splits the message to a splice of words
func MessageToWords(message string) []string {
	delimeterRule := regexp.MustCompile("[^\\w']")

	withPossibleSpace := delimeterRule.Split(message, -1)
	return removeEmptyElement(withPossibleSpace)
}

func removeEmptyElement(words []string) []string {
	var noSpace []string

	for _, word := range words {
		if word != "" {
			noSpace = append(noSpace, word)
		}
	}

	return noSpace
}

func CreateMapForMessage(words []string) map[string] {
synsetMap := make(map[string][]string)

for _, word := range words {
  synsetMap[word] = GetSynset(word)
}

return synsetMap
}
