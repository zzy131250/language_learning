package main

import "fmt"

// Guidelines for pipeline construction
// 1. stages close their outbound channels when all the send operations are done.
// 2. stages keep receiving values from inbound channels until those channels are closed or the senders are unblocked.
// Pipelines unblock senders either by ensuring there's enough buffer for all the values that are sent or
// by explicitly signalling senders when the receiver may abandon the channel.
func main() {
	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer close(done)

	in := genWithDone(done, 2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sqWithDone(done, in)
	c2 := sqWithDone(done, in)

	// Consume the first value from output.
	out := mergeWithDone(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// done will be closed by the deferred call.
}
