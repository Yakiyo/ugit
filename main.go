package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	pflag.Parse()
	pargs := pflag.Args()
	if len(pargs) < 1 {
		fmt.Fprintln(os.Stderr, "No arguments provided. Provide a command argument")
		os.Exit(1)
	}

	err := run(pargs[0], pargs[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd string, args []string) error {
	switch cmd {
	case "init":
		return initCmd()
	}

	return nil
}
