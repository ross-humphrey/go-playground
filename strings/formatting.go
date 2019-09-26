package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	// Go offers several printing "verbs" designed to format general Go Values

	fmt.Printf("%v\n", p)  // Prints out an instance
	fmt.Printf("%+v\n", p) // Variant includes the struct field names
	fmt.Printf("%#v\n", p) // Prints go syntax representation of value (source code that produces value)

	fmt.Printf("%T\n", p)    // Prints the type of a value
	fmt.Printf("%t\n", true) // Formatting a boolean

	fmt.Printf("%d\n", 123) // base 10 formatting
	fmt.Printf("%b\n", 14)  // binary representation
	fmt.Printf("%c\n", 33)  // character corresponding to the given integer
	fmt.Printf("%x\n", 456) // hex encoding

	fmt.Printf("%f\n", 78.9)        // basic decimal formatting for floats
	fmt.Printf("%e\n", 123400000.0) // scientific notation
	fmt.Printf("%E\n", 123400000.0) // slightly different scientific notation

	fmt.Printf("%s\n", "\"string\"") // basic string formatting
	fmt.Printf("%q\n", "\"string\"") // double quotes
	fmt.Printf("%x\n", "hex this")   // base-16 (hex)

	fmt.Printf("%p\n", &p) // representation of a pointer

	// To control the precision / width of a figure use a number after % in the verb
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	// Width of printed floats can be done - although better to use width.precision syntax with it
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	//To left justify use the -flag
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Control width when formatting strings
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	// and to left justify
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats and returns a string without putting it anywhere
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Format +pint to io.writers other than os.Stdout using Fprintf
	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
