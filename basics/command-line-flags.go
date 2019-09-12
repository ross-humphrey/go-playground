package main

/*
Common way to specify options for command-line programs
Go provides a flag package supporting basic command-line
flag parsing.
*/

import (
	"flag"
	"fmt"
)

func main() {

	// flag declaration
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// can also declare an option that uses existing var declared elsewhere
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var") // pointer must be parsed

	flag.Parse() // once declared - call Parse to execute the command line parsing

	// dump out parsed options
	fmt.Println("word:", *wordPtr) // not - dereference of pointers
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
