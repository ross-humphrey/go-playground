package main

import "fmt"

func main() {

	// map[key-type]val-type
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// get value
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	// remove item
	delete(m, "k2")
	fmt.Println("map:", m)

	// optional second return value when getting a map indicates if the key was present in the map
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and init a new map in one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
