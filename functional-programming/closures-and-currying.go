package main

import "fmt"

// higher order function that returns a function
// given x will return the closure with 'x' memorized on it - which can then be called (as its a function)
func add(x int) func(y int) int {

	// returns function as closure
	// variable x is obtained from outer scope of this method and memorized in the closure
	return func(y int) int {
		return x + y
	}
}

func main() {

	// we are currying the add method to create more variations
	var add10 = add(10)
	var add20 = add(20)
	var add30 = add(30)

	fmt.Println(add10(5))
	fmt.Println(add20(5))
	fmt.Println(add30(5))

}
