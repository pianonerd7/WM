package usubstitute

import (
	"strings"

	"syntacticsub/wordnet"
)

func GetSynset(word string) []string {
	resultString := wordnet.FindTheInfo_ds(word, 1, 5, 0)

	trimLeftCurly := strings.Trim(resultString, "{")
	trimRightCurly := strings.Trim(trimLeftCurly, "}")

	return strings.Split(trimRightCurly, ",")
}
