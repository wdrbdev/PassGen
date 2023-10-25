package generator

import (
	"github.com/PassGen/util"
)

// Shuffle string array to get random password
func Shuffle(chars []string, length int) string {
	var result string
	for i := 0; i < length; i++ {
		util.ShuffleArray(chars)
		result += chars[0]
	}
	return result
}
