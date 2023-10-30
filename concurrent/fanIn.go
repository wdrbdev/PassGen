package concurrent

import (
	"runtime"
)

func FanIn(generate func([]string, int) string, chars []string, length int, resultChan chan<- string, stopChan <-chan bool) {
	cpuCount := runtime.GOMAXPROCS(0)
	var stopChans []chan bool
	for i := 0; i < cpuCount; i++ {
		stopChans = append(stopChans, make(chan bool))
		go func(stopChan <-chan bool) {
			select {
			case <-stopChan:
				return
			case resultChan <- generate(chars, length):
				// do nothing, keep generating until stopped
			}
		}(stopChans[i])
	}

	go func() {
		<-stopChan
		for i := 0; i < cpuCount; i++ {
			stopChans[i] <- true
		}
	}()
}
