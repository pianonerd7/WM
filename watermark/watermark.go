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
