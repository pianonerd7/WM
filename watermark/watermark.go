package watermark

import (
	"bytes"
	"fmt"

	"code.uber.internal/engsec/syntacticsub/usubstitute"
	"code.uber.internal/engsec/syntacticsub/utility"
)

// EmbedMessage takes a message, generates a secret and watermarks the message using the secret
// and returns the message with the embedded message
func EmbedMessage(message string) string {
	synsetMap := usubstitute.GetMapFromMessage(message)
	bitSecret := utility.GetRandomBytes()
	fmt.Println(bitSecret)
	wordWithPunc := usubstitute.MessageToWords(message)

	bitIndex := 0
	var watermarkedMessage bytes.Buffer

	for _, word := range wordWithPunc {
		if synsetMap[word] != nil {
			//fmt.Printf("Index: %v, value: %v \n", bitIndex%len(bitSecret), bitSecret[bitIndex%len(bitSecret)])
			if bitSecret[bitIndex%len(bitSecret)] == utility.One {
				watermarkedMessage.WriteString(synsetMap[word][0] + " ")
			} else {
				watermarkedMessage.WriteString(word + " ")
			}
			//fmt.Printf("Word: %v, replacement: %v\n", word, synsetMap[word][0])
			bitIndex++
		} else {
			watermarkedMessage.WriteString(word + " ")
		}
	}

	if bitIndex < len(bitSecret) {
		panic("The secret was not fully embedded")
	}

	return watermarkedMessage.String()
}

// ExtractMessage takes the original message and the watermarked message
// and extracts the secret from the watermarked message
func ExtractMessage(originalMessage, embeddedMessage string) string {
	synsetMap := usubstitute.GetMapFromMessage(originalMessage)
	wordWithPuncOriginal := usubstitute.MessageToWords(originalMessage)
	wordWithPuncEmbedded := usubstitute.MessageToWords(embeddedMessage)

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

// EmbedFullMessage takes a message and embeds the message with all '1'
// it returns the watermarked message
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
