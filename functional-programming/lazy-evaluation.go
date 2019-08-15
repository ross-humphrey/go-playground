package main

import "fmt"

func addLazy(x int) int {
	fmt.Println("executing addLazy")
	return x + x
}

func multiply(x int) int {
	fmt.Println("executing multiply")
	return x * x
}

func main() {
	fmt.Println(addOrMultiply(true, addLazy, multiply, 4))
	fmt.Println(addOrMultiply(false, addLazy, multiply, 4))

}

func addOrMultiply(add bool, onAdd, onMultiply func(t int) int, t int) int {
	if add {
		return onAdd(t)
	}
	return onMultiply(t)
}
