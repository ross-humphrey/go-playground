package command_line

/**
Command line arguments are a common way to parameterize execution of programs.
*/

import (
	"fmt"
	"os"
)

func main() {

	// os.Args provides access to raw command-line arguments
	argsWithProg := os.Args
	argsWithoutProg := os.Args[:1]

	// get individual args with normal indexing
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

}
