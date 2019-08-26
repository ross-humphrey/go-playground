package main

// A panic means something unexpectedly went wrong
// Use to fail fast
import "os"

func main() {

	panic("a problem")

	// Common use is to abort if function returns error value
	// that we don't know how to (or want to) handle.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
