package utility

import (
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// getRandomBytes return a splice of bits
func getRandomBytes() []Bit {
	byteSlice := make([]byte, 2)
	rand.Read(byteSlice)
	return BytesToBits(byteSlice)
}
