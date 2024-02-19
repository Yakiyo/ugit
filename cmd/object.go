package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/ugit/data"
	"github.com/charmbracelet/log"
)

func HashObj(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid arguments provided, must provide only one argument")
	}
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
	data, err := data.GetObject(args[0], "blob")
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}
