package main

import "fmt"

func factorialTailRec(num int) int {
	return factorial(1, num)
}

func factorial(accumulator, val int) int {
	if val == 1 {
		return accumulator
	}

	return factorial(accumulator*val, val-1)
}

/**
Tail recursion in GO is benchmarked at bening slightly better than recursive - not as good as iterative,
the iterative approach should be used where performance is key.
*/
func main() {
	fmt.Println(factorialTailRec(20))
}
