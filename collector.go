package main

var WorkQueue = make(chan WorkRequest, 100)

func Collector(nn uint64) {

	work := WorkRequest{Nonce: nn}

	// Push the work onto the queue.
	WorkQueue <- work

	return
}
