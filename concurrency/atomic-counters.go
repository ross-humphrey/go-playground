package main

// atomic counters - accessed by multiple go routines

import "fmt"
import "time"
import "sync/atomic"

func main() {

	var ops uint64 // unsigned int to represent always positive counter

	for i := 0; i < 50; i++ { // start 50 go routines
		go func() {
			for {
				atomic.AddUint64(&ops, 1) // add via memory address
				// atomically increment counter
				time.Sleep(time.Millisecond) // wait between increments
			}
		}()
	}

	time.Sleep(time.Second) // wait to allow ops to accumulate

	// to safely use counter while being updated - extract copy of the current value via LoadUint64 - via memory address
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
