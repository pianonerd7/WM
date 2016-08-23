package watermark

import (
	"bytes"
	"fmt"

	"code.uber.internal/engsec/syntacticsub/sql"
	"code.uber.internal/engsec/syntacticsub/usubstitute"
	"code.uber.internal/engsec/syntacticsub/utility"
)

func EmbedMessageForAllEmail(message string, emails []string) []string {

	var embeddedMessages []string
	for _, email := range emails {
		bitSecret := utility.GetWaterMark(email)
		fmt.Println(bitSecret)
		embeddedMessages = append(embeddedMessages, EmbedMessage(message, bitSecret))
	}

	return embeddedMessages
}

// EmbedMessage takes a message, generates a secret and watermarks the message using the secret
// and returns the message with the embedded message
func EmbedMessage(message string, bitSecret []utility.Bit) string {
	synsetMap := usubstitute.GetMapFromMessage(message)
	wordWithPunc := usubstitute.MessageToWords(message)

	bitIndex := 0
	var watermarkedMessage bytes.Buffer

	for _, word := range wordWithPunc {
		if synsetMap[word] != nil {
			if bitSecret[bitIndex%len(bitSecret)] == utility.One {
				watermarkedMessage.WriteString(synsetMap[word][0] + " ")
			} else {
				watermarkedMessage.WriteString(word + " ")
			}
			bitIndex++
		} else {
			watermarkedMessage.WriteString(word + " ")
		}
	}

	if bitIndex < len(bitSecret) {
		panic("The secret was not fully embedded")
	}

	return formatEmbeddedMessage([]byte(watermarkedMessage.String()))
}

// FormatEmbeddedMessage takes a byte array and removes all the spaces
// prior to a punctuation
func formatEmbeddedMessage(embeddedMessage []byte) string {
	punctuation := "’'*?()$.,!“”"
	for index, character := range embeddedMessage {
		if bytes.IndexAny([]byte{character}, punctuation) >= 0 {
			if (index-1 >= 0) && (embeddedMessage[index-1] == ' ') {
				embeddedMessage = append(embeddedMessage[:index-1], embeddedMessage[index:]...)
			}
		}
	}
	return string(embeddedMessage[:])
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
	//return watermarkedMessage.String()
	watermark := watermarkedMessage.String()
	return sql.GetUserEmailFromWaterMark(watermark[:16])
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
