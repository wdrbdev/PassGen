package dictMarker

func Upper(dict map[string]bool) {
	var upperAlpha = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for _, char := range upperAlpha {
		dict[char] = true
	}
}
