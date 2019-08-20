package main

import "fmt"

// Specify if a channel is only meant to send or receive values
// This increases type safety of app

func ping(pings chan<- string, msg string) { // only accepts a channel for sending
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// accepts one channel to receive (pings)and second to send (pongs)
	msg := <-pings
	pongs <- msg

}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
