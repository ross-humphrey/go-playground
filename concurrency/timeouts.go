package main

// when connecting to external sources - timeouts are good!

import "time"
import "fmt"

func main() {

	c1 := make(chan string, 1) // channel created
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1: // await result -
		fmt.Println(res)
	case <-time.After(1 * time.Second): // timeout is 1 second
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second): // allow longer timeout
		fmt.Println("timeout 1")
	}

}
