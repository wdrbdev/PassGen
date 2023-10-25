package dictMarker

func Unsimilar(dict map[string]bool) {
	var similar = []string{"2", "Z", "z", "i", "I", "1", "l", "|", "L", "o", "O", "0"}
	for _, char := range similar {
		dict[char] = false
	}
}
