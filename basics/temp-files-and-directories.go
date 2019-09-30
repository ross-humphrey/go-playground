package main

// A way to create temp data (files and directories)

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// create a temp file
	f, err := ioutil.TempFile("", "sample")
	check(err)

	// display the name of the temporary file.
	fmt.Println("Temp file name:", f.Name())

	// clean up
	defer os.Remove(f.Name())

	// write some data to the file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	//if we intend to write many temp files - we can create a temp dir
	dname, err := ioutil.TempDir("", "sampledir")
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)

}
