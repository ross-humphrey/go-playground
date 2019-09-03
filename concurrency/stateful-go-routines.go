package main

/*
Use built-in synchronization features to achieve the same
result as using mutexes. (go routines and channels)

The channel-based approach aligns with Go's ideas of sharing
memory by communicating and having each piece of data owned
by exactly one goroutine
*/

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/*
In example - state will be owned by a single go routine.

To read or write that state, other goroutines will send messages
to the owning goroutine and receive corresponding replies.

These readOp and writeOp structs encapsulate those requests and a way
for goroutines to respond
*/
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// count operations
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	// go routine that owns the state
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads: // selects on the reads and writes channel, responding as they arrive
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// start 100 goroutines to issue reads to the state-owning goroutine via the reads channel
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Start 10 writes as well
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	// Take and report final operation counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

}
