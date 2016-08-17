package watermark

import (
	"bytes"
	"fmt"
	"syntacticsub/usubstitute"
	"syntacticsub/utility"
)

func EmbedMessage(message string) string {
	synsetMap := usubstitute.GetMapFromMessage(message)
	bitSecret := utility.GetRandomBytes()
	fmt.Println(bitSecret)
	wordWithPunc := usubstitute.MessageToWords(message)

	bitIndex := 0
	var watermarkedMessage bytes.Buffer

	for _, word := range wordWithPunc {
		if synsetMap[word] != nil {
			fmt.Printf("Index: %v, value: %v \n", bitIndex%len(bitSecret), bitSecret[bitIndex%len(bitSecret)])
			if bitSecret[bitIndex%len(bitSecret)] == utility.One {
				watermarkedMessage.WriteString(synsetMap[word][0] + " ")
			} else {
				watermarkedMessage.WriteString(word + " ")
			}
			fmt.Printf("Word: %v, replacement: %v\n", word, synsetMap[word][0])
			bitIndex++
		} else {
			watermarkedMessage.WriteString(word + " ")
		}
	}
	return watermarkedMessage.String()
}

func ExtractMessage(originalMessage, embeddedMessage string) string {
	synsetMap := usubstitute.GetMapFromMessage(originalMessage)
	wordWithPuncOriginal := usubstitute.MessageToWords(originalMessage)
	wordWithPuncEmbedded := usubstitute.MessageToWords(embeddedMessage)

	for i, word := range wordWithPuncEmbedded {
		fmt.Printf("(%v, %v), ", i, word)
	}

	var watermarkedMessage bytes.Buffer

	for index, word := range wordWithPuncOriginal {
		if synsetMap[word] != nil {
			if word != wordWithPuncEmbedded[index] {
				watermarkedMessage.WriteString("1")
			} else {
				watermarkedMessage.WriteString("0")
			}
		}
	}
	return watermarkedMessage.String()
}

func EmbedFullMessage(message string) string {
	synsetMap := usubstitute.GetMapFromMessage(message)
	wordWithPunc := usubstitute.MessageToWords(message)

	var watermarkedMessage bytes.Buffer

	for _, word := range wordWithPunc {
		if synsetMap[word] != nil {
			watermarkedMessage.WriteString(synsetMap[word][0] + " ")
		} else {
			watermarkedMessage.WriteString(word + " ")
		}
	}

	return watermarkedMessage.String()
}
