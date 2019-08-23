package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)
	// timers represent a single event in the future
	// tell timer how long you want to wait and channel that will be notified

	<-timer1.C // blocks on channel c until it sends a value indicating expiration
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() { // example of cancelling timer before expiration
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
