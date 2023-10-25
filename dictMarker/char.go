package dictMarker

func Char(dict map[string]bool) {
	var chars = []string{"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "_", "`", "{", "|", "}", "~"}
	for _, char := range chars {
		dict[char] = true
	}
}
