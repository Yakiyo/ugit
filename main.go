package main

import (
	"fmt"
	"os"

	"github.com/Yakiyo/ugit/cmd"
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

func run(command string, args []string) error {
	switch command {
	case "init":
		return cmd.Init()
	case "hash-object":
		return cmd.HashObj(args)
	case "cat-file":
		return cmd.CatFile(args)
	}
	// prolly some unknown command received
	return fmt.Errorf("unknown command `%v` received", command)
}
