package dictMarker

func Digit(dict map[string]bool) {
	var digits = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	for _, char := range digits {
		dict[char] = true
	}
}
