package concurrent

import (
	"runtime"
)

func FanIn(generate func([]string, int) string, chars []string, length int, resultChan chan<- string, stopChan <-chan bool) {
	cpuCount := runtime.GOMAXPROCS(0)
	for i := 0; i < cpuCount; i++ {
		go func() {
			select {
			case <-stopChan:
				return
			case resultChan <- generate(chars, length):
				// do nothing, keep generating until stopped
			}
		}()
	}
}
