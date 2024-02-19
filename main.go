package main

import (
	"fmt"
	"os"

	"github.com/Yakiyo/ugit/cmd"
	"github.com/charmbracelet/log"
	"github.com/spf13/pflag"
)

func main() {
	// dont print time in log
	log.SetTimeFormat("")
	log.SetLevel(log.InfoLevel)

	pflag.Parse()
	pargs := pflag.Args()
	if len(pargs) < 1 {
		log.Error("No arguments provided. Provide a command argument")
		os.Exit(1)
	}

	err := run(pargs[0], pargs[1:])
	if err != nil {
		log.Error(err)
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
	case "write-tree":
		return cmd.WriteTree(args)
	}
	// prolly some unknown command received
	return fmt.Errorf("unknown command `%v` received", command)
}
