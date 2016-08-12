// Package bit provides convenient way to operate bit array.
package utility

// Bit represents 1 bit data.
type Bit int

const (
	// Zero represents bit 0.
	Zero Bit = 0
	// One represents bit 1.
	One Bit = 1
)

// 1 byte consists of 8 bits.
const byteLength = 8

// BytesToBits converts byte slice to bit slice.
func BytesToBits(a []byte) []Bit {
	s := make([]Bit, len(a)*byteLength)
	for i, b := range a {
		for j := 0; j < byteLength; j++ {
			s[i*byteLength+j] = Bit((b >> uint(byteLength-1-j)) & 1)
		}
	}
	return s
}

// EqualBits returns a boolean reporting whether a and b are the same length and contain the same bits.
// A nil argument is equivalent to an empty slice.
func EqualBits(a, b []Bit) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// RotateBits returns bit slice that rotates from the given start position with the given length.
func RotateBits(a []Bit, start, length int) []Bit {
	b := make([]Bit, length)
	for i := 0; i < length; i++ {
		b[i] = a[(start+i)%len(a)]
	}
	return b
}
