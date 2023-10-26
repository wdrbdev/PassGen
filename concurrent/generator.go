package concurrent

func Generator(generate func([]string, int) string, chars []string, length int, resultCount int, resultChan chan<- string) {
	generatorChan := make(chan string, resultCount)
	go func() {
		defer close(generatorChan)
		for i := 0; i < resultCount; i++ {
			generatorChan <- generate(chars, length)
		}
	}()

	go func() {
		for result := range generatorChan {
			resultChan <- result
		}
	}()
}
