package main

import "fmt"

func main() {
	var list = []string{"Orange", "Apple", "Banana", "Grape"}
	// we are passing the array and a function as arguments to mapForEach method.
	var out = mapForEach(list, func(it string) int {
		return len(it)
	})
	fmt.Println(out)
}

func mapForEach(arr []string, fn func(it string) int) []int {
	var newArray = []int{}
	for _, it := range arr {
		// Execute the method parsed
		newArray = append(newArray, fn(it))
	}
	return newArray
}
