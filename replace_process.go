package main

import (
	"os"
	"os/exec"
	"syscall"
)

// it just replace current program process with command.

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env) // if exec is successful, the execution of the current process will be replaced by the new process
	if execErr != nil {
		panic(execErr)
	}
}
