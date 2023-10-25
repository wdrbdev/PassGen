package util

import "math/rand"

func ShuffleArray[T any](array []T) {
	for i := 0; i < len(array); i++ {
		randIdx := rand.Intn(len(array))
		array[i], array[randIdx] = array[randIdx], array[i]
	}
}
