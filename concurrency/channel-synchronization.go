package main

import "fmt"
import "time"

// Channels to sync execution across go routines.

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true // send notify we are done
}

func main() {

	done := make(chan bool, 1) // start worker go routine, giving it channel to notify
	go worker(done)

	<-done // Block until a notification from worker on channel is given
}
