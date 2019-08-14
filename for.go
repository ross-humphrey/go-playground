package main

import "fmt"

func main() {

	// Single condition loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic initial, condition, alter loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// No condition - must break the loop
	for {
		fmt.Println("loop")
		break
	}

	// Use of continue is also permitted
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	
}
