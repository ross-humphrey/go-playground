package main

/**
A line filter is a common type of program that reads input of stdin,
processes it, and then prints some derived result to stdout.grep and sed
are common line filters.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// wrapping the unbuffered os.Stdin with a buffered scanner lets us use scan

	for scanner.Scan() {
		// Text returns the current token, here the next line, from the input
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl) // write out upper cased line
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
