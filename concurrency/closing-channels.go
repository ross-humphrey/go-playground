package main

import "fmt"

// Helps communicate completion to the channel's receivers

func main() {

	jobs := make(chan int, 5) //work to be done from main()
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // receives value from jobs
			if more {         // more is fales if jobs is closed and all vals received
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done // await work using synchronization
}
