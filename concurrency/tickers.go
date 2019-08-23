package main

import "time"
import "fmt"

// Tickers are used when you want to do something repeatedly at regular intervals
func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	// Channel that is sent values.
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
