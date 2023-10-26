package concurrent

import (
	"runtime"
)

func FanIn(generate func([]string, int) string, chars []string, length int, resultCount int, resultChan chan<- string) {
	cpuCount := runtime.GOMAXPROCS(0)
	for i := 0; i < cpuCount; i++ {
		go func() {
			select {
			case resultChan <- generate(chars, length):
				// nothing to do
			default:
				// resultChan full means generation of # of password given is completed
				return
			}
		}()
	}
}
