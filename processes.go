package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(" > date")
	fmt.Println(string(dateOut))

	_, err = exec.Command("date", "-x").Output() // this will fail because -x is not a valid option
	if err != nil {
		switch e := err.(type) { // we get the error type
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "-E", "hello|goodbye")

	grepIn, _ := grepCmd.StdinPipe()   // we get the stdin pipe
	grepOut, _ := grepCmd.StdoutPipe() // we get the stdout pipe
	grepCmd.Start()                    // we have to start the command before we can write to it
	grepIn.Write([]byte("helloaa grep\ngoodbye grep"))
	grepIn.Close() // we need to close the stdin pipe
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait() // we need to wait for the read to finish
	fmt.Println("> grep hello and goodbye")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> ls -a -l -h ")
	fmt.Println(string(lsOut))

}
