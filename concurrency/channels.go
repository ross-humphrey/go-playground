package main

import "fmt"

// Channels are the pipes that connect concurrent
func main() {

	messages := make(chan string) // new channel - channels are typed

	go func() { messages <- "ping" }() // send value into channel using <-
	// <- channel syntax receives value from the channel

	msg := <-messages
	fmt.Println(msg)

	// sends and receives block until both the sender and receiver are ready.
	// Property allows us to wait at end for the "ping" without having to use synchronization

}
