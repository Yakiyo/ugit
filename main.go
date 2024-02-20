package main

import (
	"fmt"
	"os"

	"github.com/Yakiyo/ugit/cmd"
	"github.com/Yakiyo/ugit/utils"
	"github.com/charmbracelet/log"
)

func main() {
	// dont print time in log
	log.SetTimeFormat("")
	log.SetLevel(log.InfoLevel)
	args := os.Args[1:]

	if len(args) < 1 {
		log.Error("No arguments provided. Provide a command argument")
		os.Exit(1)
	}

	err := run(args[0], args[1:])
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func run(command string, args []string) error {
	// regular commands go here, the ones that dont
	// require any initializations
	switch command {
	case "init":
		return cmd.Init()
	}

	// this commands must only be ran in a repo
	// which is essentially a dir with a .ugit dir in it
	utils.CwdIsRepo()
	switch command {
	case "hash-object":
		return cmd.HashObject(args)
	case "cat-file":
		return cmd.CatFile(args)
	case "write-tree":
		return cmd.WriteTree(args)
	case "read-tree":
		return cmd.ReadTree(args)
	}
	// prolly some unknown command received
	return fmt.Errorf("unknown command `%v` received", command)
}
