package main

import "fmt"
import "sort"

// go sort package implements sorting for builtins and user defined types
func main() {

	strs := []string{"c", "a", ",b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	s := sort.IntsAreSorted(ints) // checks if sorted
	fmt.Println("Sorted: ", s)
}
