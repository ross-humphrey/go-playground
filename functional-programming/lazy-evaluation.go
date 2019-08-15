package main

import "fmt"

func add(x int) int {
	fmt.Println("executing add")
	return x + x
}

func multiply(x int) int {
	fmt.Println("executing multiply")
	return x * x
}

func main() {
	fmt.Println(addOrMultiply(true, add, multiply, 4))
	fmt.Println(addOrMultiply(false, add, multiply, 4))

}

func addOrMultiply(add bool, onAdd, onMultiply func(t int) int, t int) int {
	if add {
		return onAdd(t)
	}
	return onMultiply(t)
}
