package main

var WorkQueue = make(chan WorkRequest, 100)

func Collector(url string) {

	work := WorkRequest{Url: url}

	// Push the work onto the queue.
	WorkQueue <- work

	return
}
