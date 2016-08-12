package main

import (
	"fmt"
	"syntacticsub/usubstitute"
	"syntacticsub/wordnet"
)

func main() {
	wordnet.InitWN()
	word := "happy"
	//fmt.Printf("Sense Results for '%s': %s\n", word, wordnet.FindTheInfo(word, 1, 5, 0))
	//fmt.Printf("Sense Results for '%s': %s\n", word, wordnet.FindTheInfo_ds(word, 1, 5, 0))
	fmt.Println(usubstitute.GetSynset(word))
}
