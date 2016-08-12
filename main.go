package main

import (
	"fmt"
	"syntacticsub/wordnet"
)

func main() {
	wordnet.InitWN()
	word := "root"
	fmt.Printf("Sense Results for '%s': %s\n", word, wordnet.FindTheInfo(word, 1, 5, 0))
}
