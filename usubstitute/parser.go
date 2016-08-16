package usubstitute

import (
	"fmt"
	"regexp"
	"strings"

	"syntacticsub/sql"
	"syntacticsub/wordnet"
)


const (
	"Noun" = {"EX", "NN", "NNP", "NNPS", "NNS", "NNS|VBZ", "NN|CD", "NN|JJ"},
	Verb = {"VB", "VBD", "VBG", "VBG|NN", "VBN", "VBP", "VBZ"},
	Adjective = {"JJ", "JJR", "JJS"},
	Adverb = {"EX", "WRB", "RP", "RB", "RBR", "RBS"},
)

PartsOfSpeechTagMap := map[string]string {
	"EX" : "Noun",
	"NN" : "Noun",
	"NNP" : "Noun",
	"NNPS" : "Noun",
	"NNS" : "Noun",
	"NNS|VBZ" : "Noun",
	"NN|CD" : "Noun",
	"NN|JJ" : "Noun",
	"VB" : "Verb",
	"VBD" : "Verb",
	"VBG" : "Verb",
	"VBG|NN" : "Verb",
	"VBN" : "Verb",
	"VBP" : "Verb",
	"VBZ" : "Verb",
	"JJ" : "Adjective",
	"JJR" : "Adjective",
	"JJS" : "Adjective",
	"EX" : "Adverb",
	"WRB" : "Adverb",
	"RP" : "Adverb",
	"RB" : "Adverb",
	"RBR" : "Adverb",
	"RBS" : "Adverb",
}

func GetPOS(word string) string {
	words := sql.QueryByWord(word)

	var pos string
	if len(words) > 0 {
		pos = words[0].POS
		fmt.Println(pos)
	}


	return ""
}

// getSynset takes in a word and returns a synset
func getSynset(word string) []string {
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

// messageToWords takes in a string of words representing a message
// and splits the message to a splice of words
func MessageToWords(message string) []string {
	fmt.Println(message)
	delimeterRule := regexp.MustCompile(`[A-Za-z’]+|[*?()$.,!“”–]`)

	withPossibleSpace := delimeterRule.FindAllString(message, -1)
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

// createMapForMessage takes in a slice of words and finds the synset
// for every word in the slice, and if the synset is not empty, it
// adds the word synset pair to the map
func createMapForMessage(words []string) map[string][]string {
	synsetMap := make(map[string][]string)

	for _, word := range words {
		synset := getSynset(word)
		if synset != nil {
			synsetMap[word] = synset
		}
	}
	return synsetMap
}

func GetMapFromMessage(message string) map[string][]string {
	words := MessageToWords(strings.ToLower(message))
	return createMapForMessage(words)
}

type PuncLoc struct {
	SliceIndex  int
	Punctuation string
}

// Is this function really needed?
func getPunctuationIndex(words []string) []PuncLoc {
	var punctuationIndex []PuncLoc

	delimeterRule := regexp.MustCompile(`[*?()$.,!“”–]`)

	for index, word := range words {
		isPunctuation := delimeterRule.FindAllString(word, -1)
		if len(isPunctuation) == 1 {
			newPuncLoc := PuncLoc{
				SliceIndex:  index,
				Punctuation: word,
			}
			punctuationIndex = append(punctuationIndex, newPuncLoc)
		}
	}
	fmt.Println(punctuationIndex)
	return punctuationIndex
}
