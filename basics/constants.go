package main

import "fmt"
import "math"

const s string = "constant"

// working with constants
func main() {
	fmt.Println(s)

	// n is not assigned a type here
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	// but n is given a type in context
	fmt.Println(math.Sin(n))

}
