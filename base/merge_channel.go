package base

import (
	"sync"
)

// Make Merge Channel
func MakeMergeChannel[T any](upstreams ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup

	// Start an output goroutine for each input channel in upstreams.
	wg.Add(len(upstreams))
	for _, c := range upstreams {
		go func(c <-chan T) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}

	// Start a goroutine to close out once all the output goroutines are done.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
