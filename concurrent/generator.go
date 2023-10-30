package concurrent

func Generator(generate func([]string, int) string, chars []string, length int, resultChan chan<- string, stopChan <-chan bool) {
	generatorChan := make(chan string, 16)
	go func() {
		for {
			select {
			case <-stopChan:
				close(generatorChan)
				return
			case generatorChan <- generate(chars, length):
				// do nothing, keep generating until stopped
			}

		}
	}()

	go func() {
		for result := range generatorChan {
			resultChan <- result
		}
	}()
}
