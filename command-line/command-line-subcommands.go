package command_line

/**
subcommands - each have their own set of flags - like go build and go get
The flag package lets us easily define simple subcommands that have
their own flags
*/

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// declare a subcommand using the NewFlagSet function and proceed to define new flags specific for subcommand
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// For different subcommand we can define different supported flags
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Subcommand is expected as the first argument to the program
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// check which subcommand is invoked
	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println(" enable:", *fooEnable)
		fmt.Println(" name:", *fooName)
		fmt.Println(" tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println(" level:", *barLevel)
		fmt.Println(" tail", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

}

// call using:
// ./command-line-subcommands foo -enable -name=joe a1 a2
// ./command-line-subcommands bar -level 8 a1
// ./command-line-subcommands bar -enable a1
