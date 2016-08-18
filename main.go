package main

import (
	"fmt"

	"code.uber.internal/engsec/syntacticsub/watermark"
	"code.uber.internal/engsec/syntacticsub/wordnet"
)

func main() {
	wordnet.InitWN()
	//010100101111110101010010111111010101001011111101 should be
	//010100101111110101010010111111010101       actual
	//[0 1 0 1 0 0 1 0 1 1 1 1 1 1 0 1]
	word := "It was a big, bold idea, especially given that Uber was still a relatively small and no one in China had ever even heard of us.  And of course, anytime we got into a discussion about our efforts in China, most people thought we were naive, crazy – or both. We saw things differently of course. China is an inspiring country with astonishing opportunity. Many of the world’s mega cities are Chinese, and their thirst for transportation innovation is second to none. Uber’s mission to make “transportation as reliable as running water, everywhere for everyone” resonates especially strongly in China."
	embeded := "It was a big , bluff idea , particularly given that Uber was still a relatively little and no one in China had always yet heard of us . And of class , anytime we got into a treatment about our efforts in PRC , most citizenry thought we were naive , brainsick or both . We saw things differently of class . China is an inspiring state with astonishing opportunity . many of the world’s mega cities are Chinese , and their thirstiness for transit invention is second to none . Uber’s charge to do “ transit as reliable as running urine , everywhere for everyone ” resonates particularly strongly in PRC ."
	//fmt.Println(wordnet.GetSenseLength("happily", 3, 23))
	//fmt.Println(wordnet.GetSenseLength("blue", 3, 23))
	//fmt.Printf("Sense Results for '%s': %s\n", "blue", wordnet.FindTheInfo_ds("blue", 3, 23, 0))
	//fmt.Printf("Sense Results for '%s': %s\n", "blue", wordnet.FindTheInfo("blue", 3, 2, 0))
	//fmt.Printf("Sense Results for '%s': %s\n", "blue", wordnet.FindTheInfo("blue", 3, 3, 0))

	//fmt.Printf("Sense Results for '%s': %s\n", "blue", wordnet.FindTheInfo_ds("blue", 1, 5, 0))
	/*x := usubstitute.GetMapFromMessage(word)
	for key, value := range x {
		fmt.Printf("%v , %v \n", key, value)
	}*/
	//fmt.Println(watermark.EmbedMessage(word))
	fmt.Println(watermark.ExtractMessage(word, embeded))
	//watermark.EmbedMessage(word)
	//watermark.ExtractMessage(word, embeded)
	//usubstitute.GetMapFromMessage(word)
}
