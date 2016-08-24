package main

import (
	"fmt"
	"os"
	"strconv"

	"code.uber.internal/engsec/syntacticsub/utility"
	"code.uber.internal/engsec/syntacticsub/watermark"
	"code.uber.internal/engsec/syntacticsub/wordnet"
)

func main() {
	wordnet.InitWN()

	// [1, 2, 3, 4]
	// original message, watermarked message, recipient, command
	args := os.Args[1:]
	originalMessage := args[0]
	watermarkedMessage := args[1]
	recipient := utility.SplitEmailToSlice(args[2])
	command, _ := strconv.Atoi(args[3])
	switch command {
	case 1:
		fmt.Printf(watermark.EmbedFullMessage(originalMessage))
	case 2:
		fmt.Printf("%v", watermark.EmbedMessageForAllEmail(originalMessage, recipient))
	case 3:
		fmt.Printf(watermark.ExtractMessage(originalMessage, watermarkedMessage))
	}
}
