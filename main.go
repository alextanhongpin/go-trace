package main

import (
	"context"
	"os"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ctx := context.Background()

	// Start a region.
	reg := trace.StartRegion(ctx, "foo")

	// Log custom metrics.
	trace.Log(ctx, "bar", "baz")

	// End the region.
	reg.End()

	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()
	// read from channel
	<-ch
}
