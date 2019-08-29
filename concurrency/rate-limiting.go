package main

import "time"
import "fmt"

// Rate limiting is a mechanism for controlling resource utiliziation and QOS

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// The limiter channel will receive a value every 200 ms. Regulator in our scheme
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Bursty limiter - allow short bursts of requests
	burstyLimiter := make(chan time.Time, 3)

	// fill up channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// every 200 ms we will add a new value to the bursty limiter
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 incoming requests.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
