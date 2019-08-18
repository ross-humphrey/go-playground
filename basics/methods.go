package main

import "fmt"

type rect struct {
	width, height int
}

// Area method has a receiver type of *rect (pointer shown below)
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types (value shown below)
func (r rect) perim() int {
	return 2 * r.width * 2 * r.height
}

func main() {

	r := rect{10, 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for method valls
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim", rp.perim())
}
