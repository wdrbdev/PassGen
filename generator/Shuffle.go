package generator

import (
	"math/rand"
)

// Shuffle string array to get random password
func Shuffle(chars []string, length int) string {
	var result string
	for i := 0; i < length; i++ {
		shuffleArray(chars)
		result += chars[0]
	}
	return result
}

func shuffleArray[T any](array []T) {
	for i := 0; i < len(array); i++ {
		randIdx := rand.Intn(len(array))
		temp := array[i]
		array[i] = array[randIdx]
		array[randIdx] = temp
	}
}
