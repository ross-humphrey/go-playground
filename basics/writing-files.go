package main

// writing files in GO follows similar patterns to the ones we saw earlier for reading

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// how to dump a string into a file
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// for more granular writes, open a file for writing
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	// Write byte slices
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes \n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes \n", n3)

	// Issue a Sync to flush writes to stable storage
	f.Sync()

	// buffered writers in addition to buffered readers
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// use Flush to ensure all buffered operations have been applied to the underlying writer
	w.Flush()

}
