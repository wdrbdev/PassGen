package generator

import (
	"math/rand"
)

// Shuffle string array to get random password
func RandIdx(chars []string, length int) string {
	var result string
	for i := 0; i < length; i++ {
		randIdx := rand.Intn(len(chars))
		result += chars[randIdx]
	}
	return result
}
