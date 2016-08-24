package utility

import (
	"math/rand"
	"strings"
	"time"

	"code.uber.internal/engsec/syntacticsub/sql"
)

func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// getRandomBytes return a splice of bits
func GetRandomBytes() ([]Bit, string) {
	var bits []Bit
	var stringBits string
	count := -1
	for count != 0 {
		bits = getRandomByte()
		stringBits = ToString(bits)
		count = sql.CountUserWatermark(stringBits)
	}
	return bits, stringBits
}

func getRandomByte() []Bit {
	byteSlice := make([]byte, 2)
	rand.Seed(time.Now().Unix())
	rand.Read(byteSlice)
	return BytesToBits(byteSlice)
}

// GetWaterMark takes an email and checks to see if this email
// is already in the database. If not, it generates a unique
// watermark and stores it into the database
func GetWaterMark(email string) []Bit {

	existingWaterMark := sql.GetUserWatermarkFromEmail(email)
	if existingWaterMark == "" {
		bits, StringBit := GetRandomBytes()
		sql.InsertNewUser(email, StringBit)
		return bits
	}
	return ToBitSlice(existingWaterMark)
}

func SplitEmailToSlice(emails string) []string {
	return strings.Split(emails, ",")
}
