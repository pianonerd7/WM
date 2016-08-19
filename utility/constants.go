package utility

//func GetUnicodeASCIIMap() map[string]string {
//}

// GetPOSMap returns a map (key, value) <--> (POS, POS tag in WordNet)
func GetPOSMap() map[string]int {
	partsOfSpeechMap := map[string]int{
		"Noun":      1,
		"Verb":      2,
		"Adjective": 3,
		"Adverb":    4,
	}
	return partsOfSpeechMap
}

// GetPOSTagMap returns a map (key, value) <--> (POS tag in ANC, POS)
func GetPOSTagMap() map[string]string {
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
