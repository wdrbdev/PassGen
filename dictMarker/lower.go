package dictMarker

func Lower(dict map[string]bool) {
	var upperAlpha = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, char := range upperAlpha {
		dict[char] = true
	}
}
