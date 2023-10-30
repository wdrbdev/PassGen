package concurrent

import (
	"runtime"
)

type job struct {
	generate func([]string, int) string
	chars    []string
	length   int
}

func WorkerPool(generate func([]string, int) string, chars []string, length int, resultChan chan<- string, stopChan <-chan bool) {
	workerCount := runtime.GOMAXPROCS(0)
	jobChan := make(chan job, 16)
	var workerStopChans []chan bool
	for i := 0; i < workerCount; i++ {
		workerStopChan := make(chan bool)
		workerStopChans = append(workerStopChans, workerStopChan)

		go worker(jobChan, resultChan, workerStopChans[i])
	}

	go func() {
		for {
			// fan-out stop signal or jobs
			select {
			case <-stopChan:
				for i := 0; i < workerCount; i++ {
					workerStopChans[i] <- true
				}
				return
			case jobChan <- job{generate: generate, chars: chars, length: length}:
				// do nothing, keep generating until stopped
			}
		}
	}()
}

func worker(jobChan <-chan job, resultChan chan<- string, workerStopChan <-chan bool) {
	for {
		select {
		case <-workerStopChan:
			return
		case job := <-jobChan:
			// fan-in result
			resultChan <- job.generate(job.chars, job.length)
		}
	}
}
