package main

// select with default clause - non blocking send and receives
// non blocking multi way selects

import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	// non blocking receive, if value available - select
	// takes the <- messages case with value, if not defaults
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no messages received")
	}

	// non blocking send
	msg := "hi"
	select {
	case messages <- msg: // no buffer can't be sent so default
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message sent")
	}

	// implement a multi-way non blocking select
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
