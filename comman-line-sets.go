package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)   // flag set defines a new set of flags for any run of the program
	fooEnable := fooCmd.Bool("enable", false, "boolean") // flags are defined relative to the flag set
	fooName := fooCmd.String("name", "", "string")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError) // exit on error if the flag is not defined
	barLevel := barCmd.Int("level", 0, "int")

	if len(os.Args) < 2 { // means that the program was run without any arguments
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1) // exit with status 1 (unsuccessful termination)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println(" enable:", *fooEnable)
		fmt.Println(" name:", *fooName)
		fmt.Println(" tail:", fooCmd.Args()) // returns the non-flag arguments
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel) // level could be used without -level flag but with just level
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
