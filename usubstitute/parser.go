package usubstitute

import (
	"fmt"
	"regexp"
	"strings"

	"code.uber.internal/engsec/syntacticsub/sql"
	"code.uber.internal/engsec/syntacticsub/wordnet"
)

func getPOSTagMap() map[string]string {
	partsOfSpeechTagMap := map[string]string{
		"NN":      "Noun",
		"NNP":     "Noun",
		"NNPS":    "Noun",
		"NNS":     "Noun",
		"NNS|VBZ": "Noun",
		"NN|CD":   "Noun",
		"NN|JJ":   "Noun",
		"VB":      "Verb",
		"VBD":     "Verb",
		"VBG":     "Verb",
		"VBG|NN":  "Verb",
		"VBN":     "Verb",
		"VBP":     "Verb",
		"VBZ":     "Verb",
		"JJ":      "Adjective",
		"JJR":     "Adjective",
		"JJS":     "Adjective",
		"WRB":     "Adverb",
		"RP":      "Adverb",
		"RB":      "Adverb",
		"RBR":     "Adverb",
		"RBS":     "Adverb",
	}
	return partsOfSpeechTagMap
}

func getPOS(word string) string {
	words := sql.QueryByWord(word)

	var pos string
	if len(words) > 0 {
		pos = words[0].POS
	}

	//fmt.Println("POS TAG IS...")
	//fmt.Println(getPOSTagMap()[pos])
	return getPOSTagMap()[pos]
}

// getSynset takes in a word and returns a synset
func getSynset(word string, wordnetPOSNumber, whichSense int) []string {
	resultString := wordnet.FindTheInfo_ds(word, wordnetPOSNumber, 23, whichSense)

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

func GetAllSynset(word string) []string {
	PossiblePOSTag := getPOS(word)
	wordnetPOSNumber := wordnet.GetPOSMap()[PossiblePOSTag]
	//fmt.Println(wordnetPOSNumber)

	if wordnetPOSNumber == 0 {
		return nil
	}

	senseCount := wordnet.GetSenseLength(word, wordnetPOSNumber, 23)
	var uniqueSenses []string

	for i := 1; i <= senseCount; i++ {
		newSet := getSynset(word, wordnetPOSNumber, i)
		for _, newSense := range newSet {
			if !didPassAllFilter(newSense, uniqueSenses) {
				uniqueSenses = append(uniqueSenses, newSense)
			}
		}
	}

	if len(uniqueSenses) < 2 {
		return uniqueSenses
	}

	return sortStringSliceByFrequency(uniqueSenses)
}

func sortStringSliceByFrequency(words []string) []string {
	var wordSlice sql.Words

	for _, word := range words {
		//fmt.Println(word)
		newWord := sql.GetHighestFreqForWord(word)

		x := newWord.Word
		if x != "" {
			wordSlice = append(wordSlice, &newWord)
		}
	}

	return wordsSliceToStringSlice(wordSlice)
}

func wordsSliceToStringSlice(words sql.Words) []string {
	/*fmt.Println()
	for _, o := range words {
		fmt.Println(o)
	}
	*/
	var stringSlice []string

	for _, word := range words {
		if word.Word != "" {
			stringSlice = append(stringSlice, word.Word)
		}
	}

	return stringSlice
}

// messageToWords takes in a string of words representing a message
// and splits the message to a splice of words
func MessageToWords(message string) []string {
	delimeterRule := regexp.MustCompile(`[A-Za-z-]+|[A-Za-z’]+|[*?()$.,!“”]`)
	//delimeterRule := regexp.MustCompile(`[A-Za-z’]+|[*?()$.,!“”]`)

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
		synset := GetAllSynset(word)
		if synset != nil {
			synsetMap[word] = synset
		}
	}
	return synsetMap
}

func GetMapFromMessage(message string) map[string][]string {
	words := MessageToWords(message)
	fmt.Println(words)
	return createMapForMessage(words)
}
