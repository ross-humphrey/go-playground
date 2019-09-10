package main

/**
Env variables convey configuration information to Unix programs
This shows how to get, set and list env vars
*/

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Set
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	// empty string if not present
	fmt.Println("BAR", os.Getenv("BAR"))

	// use os.Environ to list all key/val pairs
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}

}
