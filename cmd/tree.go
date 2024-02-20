// commands related to creating & reading
//  tree objects / directories

package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/ugit/base"
	"github.com/Yakiyo/ugit/utils"
	"github.com/charmbracelet/log"
)

// scan cwd and create objects of the entire dir recursively
func WriteTree(args []string) error {
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

// read an existing tree object and write it to current directory
func ReadTree(args []string) error {
	utils.NArgs(args, 1)
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	return base.ReadTree(args[0], cwd)
}
