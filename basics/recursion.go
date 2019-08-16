package main

import "fmt"

func fact(n int) int {
	if n == 0 { // base case
		return 1
	}
	return n * fact(n-1) // calls its self until it reaches base case
}

func main() {
	fmt.Println(fact(7))
}
