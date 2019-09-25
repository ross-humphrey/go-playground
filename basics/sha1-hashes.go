package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	s := "sha1 this string"

	h := sha1.New()

	// Pattern for generating a hash is sha1.New(), write then sum.
	h.Write([]byte(s)) // use byte(s) to coerce into a bytes
	bs := h.Sum(nil)   // get final hash as a byte slice
	fmt.Println(s)
	fmt.Println("%x\n", bs) //use %x format verb to convert hash results to a hex string
}
