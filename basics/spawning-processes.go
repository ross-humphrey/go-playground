package main

/**
Go programs sometimes need to spawn other non-Go process
*/

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	dateCmd := exec.Command("date")

	/*
		.Output is another helper that handles the common case of running a command,
		waiting for it to finish, collecting the output
	*/

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Next - an example where we pipe data to external process on stdin
	grepCmd := exec.Command("grep", "hello")

	// Grab input output pipes, start the process and write some input to it, read the output and wait for exit
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbyegrep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	/**
	When spawning commands we need to provide an explicitly delineated command and argument array vs
	being able to pass in one cmd line string. To spawn a full command with a string, use bash's -c option
	*/

	lsCmd := exec.Command("bash", "-c", "-ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}
