package main

import "fmt"

// range iterates over elements in a variety of data structures
func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums { // range on arrays and slices provides both index and value for each entry
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on a map iterates over key-value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can iterate over only keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings - iterates over unicode code points. The first value is the starting byte index of rune,
	// the second is the run itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
