package main

import "code.uber.internal/engsec/syntacticsub/wordnet"

func main() {
	wordnet.InitWN()
	//010100101111110101010010111111010101001011111101 should be
	//0101001011111101010100            actual
	//[0 1 0 1 0 0 1 0 1 1 1 1 1 1 0 1]
	//[0 1 0 1 0 0 1 0 1 1 1 1 1 1 0 1]
	//word := "Being an entrepreneur means you are an explorer by nature, doing what everyone thinks is impossible but with an optimistic perspective on the unknown. Uber entered this uncharted territory in February 2014, two years after Didi was founded. We were a young American business entering a country where most US internet companies had failed to crack the code, and with a product that needed rebuilding. Our China effort has been one of Uber's most entrepreneurial because we literally had to start from scratch."
	//embeded := "Being an entrepreneur means you are an adventurer by nature, doing what everyone thinks is impossible but with an affirmative perspective on the unknown. Uber entered this chartless territory in Feb, ii age after Didi was founded. We were a immature American concern entering a state where most US internet companies had failed to cleft the code, and with a merchandise that needed rebuilding. Our China attempt has been one of Uber's most entrepreneurial because we literally had to start from scratch."
	//fmt.Println(watermark.EmbedMessage(word))
	//fmt.Println(watermark.ExtractMessage(word, embeded))
	//watermark.EmbedMessage(word)
	//watermark.ExtractMessage(word, embeded)

	/*
		args := os.Args[1:]
		text := args[0]
		command, _ := strconv.Atoi(args[1])
		switch command {
		case 1:
			fmt.Printf(watermark.EmbedFullMessage(text))
		case 2:
			fmt.Printf("hello world")
		}
	*/

}
