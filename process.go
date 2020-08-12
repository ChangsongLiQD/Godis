package main

type Process func()

var (
	processChan = make(chan Process)
	doneChan    = make(chan bool)
)

func init() {
	go func() {
		for proc := range processChan {
			proc()
			doneChan <- true
		}
	}()
}

func DoProcess(p Process) {
	processChan <- p
	<-doneChan
}
