package main

/*
Signals are used to handle UNIX signals such as SIGTERM gracefully
*/

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

	// go signal notification works by sending os.Signal values on a channel
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal notify registers the given channel to receive notifications of the specified signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// goroutine executes a blocking receive for signals
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
