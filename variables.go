package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	// type can be declared
	var b,c int = 1, 2
	fmt.Println(b, c)

	// type is inferred
	var d = true
	fmt.Println(d)

	// no init = zero-valued
	var e int
	fmt.Println(e)

	// := short hand for declaring and init of var
	f := "apple"
	fmt.Println(f)
}
