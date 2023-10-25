package generator

import (
	"github.com/PassGen/util"
)

// Shuffle string array to get random password
func RandIdxDeduplicate(chars []string, length int) string {
	var result string
	util.ShuffleArray(chars)
	for i := 0; i < length; i++ {
		result += chars[i]
	}
	return result
}
