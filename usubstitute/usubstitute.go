package usubstitute

import (
	"regexp"
	"strings"

	"syntacticsub/wordnet"
)

//TODO: have constants for parts of speech and integrate with ANC and BNC
const ()

// GetSynset takes in a word and returns a synset
func GetSynset(word string) []string {
	resultString := wordnet.FindTheInfo_ds(word, 1, 5, 0)

	if resultString == "" {
		return nil
	}

	delimeterRules := func(t rune) bool {
		return t == '{' || t == '}' || t == ',' || t == ' '
	}

	mayContainOriginalWord := strings.FieldsFunc(resultString, delimeterRules)

	var wordList []string

	// Remove original word from slice
	for _, wordInList := range mayContainOriginalWord {
		if wordInList != word {
			wordList = append(wordList, wordInList)
		}
	}
	return wordList
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

func CreateMapForMessage(words []string) map[string][]string {
	synsetMap := make(map[string][]string)

	for _, word := range words {
		synset := GetSynset(word)
		if synset != nil {
			synsetMap[word] = synset
		}
	}
	return synsetMap
}

func CallEverything(message string) map[string][]string {
	words := MessageToWords(strings.ToLower(message))
	return CreateMapForMessage(words)
}
