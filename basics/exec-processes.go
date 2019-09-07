package main

import "syscall"
import "os"
import "os/exec"

/*
Use exec to replace current GO process with another
*/

func main() {

	binary, lookErr := exec.LookPath("ls") // use exec.LookPath to find where ls sits (/bin/ls)
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ() // exec needs environment variables

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	// go does not offer classic Unix fork function
}
