package dictMarker

func Char(dict map[string]bool) {
	var upperAlpha = []string{"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "_", "`", "{", "|", "}", "~"}
	for _, char := range upperAlpha {
		dict[char] = true
	}
}
