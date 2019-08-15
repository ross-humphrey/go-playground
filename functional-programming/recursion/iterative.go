package main

import "fmt"

// factorial using a traditional iterative approach
func factorialIterative(num int) int {
	result := 1
	for ; num > 0; num-- {
		result *= num
	}
	return result
}

func main() {
	fmt.Println(factorialIterative(20))

}
