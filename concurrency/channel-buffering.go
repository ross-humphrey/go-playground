package main

import "fmt"

//Buffered channels accept a limited number of values without a corresponding receiver for vals
func main() {

	messages := make(chan string, 2) // channel of strings buffering up to 2 vals

	messages <- "buffered"
	messages <- "channel" // no receiver

	fmt.Println(<-messages) // receive as usual
	fmt.Println(<-messages)

}
