package main

import "fmt"
import "os"

func main() {

	defer fmt.Println("!") // defer will not be run when using os.Exit - fmt.Println will never be called

	os.Exit(3) // use it to immediately exit with a given status
}
