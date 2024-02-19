// commands related to creating & reading
//  tree objects / directories

package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/ugit/base"
	"github.com/charmbracelet/log"
)

func WriteTree(args []string) error {
	cwdIsRepo()
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get cwd, err = %v", err)
	}
	log.Infof("cwd %v", cwd)
	id, err := base.WriteTree(cwd)
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}

func ReadTree(args []string) error {
	cwdIsRepo()
	nArgs(args, 1)
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	return base.ReadTree(args[0], cwd)
}
