package main

import (
	"fmt"
	"sync"
	"time"
)

// Wait for multiple goroutines to finish, use wait group

func workerW(id int, wg *sync.WaitGroup) { //wait group must be passed to functions by pointer
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	wg.Done() // notify wait group the worker is done
}

func main() {

	var wg sync.WaitGroup // wait group used to wait for all go routines launched here to finish

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerW(i, &wg) // launch multiple go routines and increment waitGroup counter for each
	}

	wg.Wait() // Block until all workers notify they are done

}
