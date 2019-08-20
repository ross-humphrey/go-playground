package main

// light weight threads of execution

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct") // direct synchronous

	go f("goroutine") // concurrently with the calling function

	//anonymous function call - can start go routine from there

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln() // key press before exit
	fmt.Println("done")

}
