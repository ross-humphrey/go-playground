package main

/**
Defer is used to ensure that a function call is performed later in a program's execution (usually for cleanup)
(used where finally and ensure would be used in other languages)
*/

import "fmt"
import "os"

func main() {
	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	_, _ = fmt.Fprintf(f, "data")
}

func closeFile(f *os.File) {
	fmt.Print("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error %v\n", err)
		os.Exit(1)
	}
}
