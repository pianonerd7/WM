package watermark

import (
	"bytes"
	"syntacticsub/usubstitute"
)

func EmbedMessage(message string) string {
	synsetMap := usubstitute.GetMapFromMessage(message)
	//bitSecret := utility.GetRandomBytes()
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
