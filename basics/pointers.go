package main

import "fmt"

// passed by value - zeroval gets a copy of ival distinct from one in the calling function
func zeroval(ival int) {
	ival = 0
}

// passes pointer, dereferences the pointer - assigning value will change value at ref'd address
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // address of pointer
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer", &i)
}
