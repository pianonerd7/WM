package utility

import (
	"math/rand"
	"time"
)

func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// getRandomBytes return a splice of bits
func GetRandomBytes() []Bit {
	byteSlice := make([]byte, 2)
	rand.Seed(time.Now().Unix())
	rand.Read(byteSlice)
	return BytesToBits(byteSlice)
}
