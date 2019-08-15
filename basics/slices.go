package main

import "fmt"

// Slices - data type
func main() {

	// slices are typed only by elements they contain (not number of elements)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copying slices
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5] // slice between a range
	fmt.Println("sl1:", l)

	l = s[:5] // slice up to excluding s[5]
	fmt.Println("sl2:", l)

	l = s[2:] // slice up from index s[2] - including
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} //declare and init a variable for slice in a single line
	fmt.Println("dcl:", t)

	// Slices can be composed into multi-dimensional data structures - the length of inner slices can vary,
	// unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
