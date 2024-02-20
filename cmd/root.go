// package command contains functions implementing the
// command line interface
package cmd

// this file contains general commands that don't require a separate file
// or those which are pretty short and their actual implementation lives in
// a separate file/package/dir

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Yakiyo/ugit/base"
	"github.com/Yakiyo/ugit/data"
	"github.com/Yakiyo/ugit/utils"
)

func Init() error {
	err := os.MkdirAll(data.GIT_DIR, os.ModePerm)
	if err != nil {
		return err
	}
	return os.MkdirAll(filepath.Join(data.GIT_DIR, "objects"), os.ModePerm)
}

func Commit(args []string) error {
	utils.NArgs(args, 1)
	id, err := base.Commit(args[0])
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}
