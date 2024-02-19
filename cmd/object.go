package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/ugit/base"
	"github.com/Yakiyo/ugit/data"
	"github.com/charmbracelet/log"
)

func HashObj(args []string) error {
	cwdIsRepo()
	nArgs(args, 1)
	content, err := os.ReadFile(args[0])
	if err != nil {
		log.Error("failure when reading file", "path", args[0])
		return err
	}
	h, err := data.CreateObject(content, "blob")
	if err != nil {
		return err
	}
	fmt.Println(h)
	return nil
}

func CatFile(args []string) error {
	cwdIsRepo()
	nArgs(args, 1)
	data, err := data.GetObject(args[0], "blob")
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

func WriteTree(args []string) error {
	cwdIsRepo()
	nArgs(args, 1)
	return base.WriteTree(args[0])
}